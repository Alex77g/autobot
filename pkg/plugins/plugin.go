package plugins

import (
	"context"
	"sync"
	"time"

	"github.com/andersnormal/autobot/pkg/plugins/filters"
	"github.com/andersnormal/autobot/pkg/plugins/runtime"
	pb "github.com/andersnormal/autobot/proto"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"
)

const (
	defaultRegisterTimeout = 10 * time.Second
)

// Event symbolizes events that can occur in the plugin.
// Though they are generally triggered from the sever.
type Event int32

const (
	ReadyEvent         Event = 0
	RefreshConfigEvent Event = 1
)

// Plugin describes a plugin in general.
// It should not be instantiated directly.
type Plugin struct {
	opts *Opts

	sc stan.Conn
	nc *nats.Conn

	ready          chan struct{}
	events         chan Event
	eventsChannels []chan Event
	resp           string

	env runtime.Env

	ctx     context.Context
	cancel  func()
	errOnce sync.Once
	err     error
	wg      sync.WaitGroup

	logger *log.Entry

	sync.RWMutex
}

// Opt is an Option
type Opt func(*Opts)

// Opts are the available options
type Opts struct {
	RegisterTimeout time.Duration
}

// SubscribeFunc ...
type SubscribeFunc = func(*pb.Event) (*pb.Event, error)

// WithContext is creating a new plugin and a context to run operations in routines.
// When the context is canceled, all concurrent operations are canceled.
func WithContext(ctx context.Context, env runtime.Env, opts ...Opt) (*Plugin, context.Context) {
	p := newPlugin(opts...)

	ctx, cancel := context.WithCancel(ctx)
	p.cancel = cancel
	p.ctx = ctx
	p.env = env

	return p, ctx
}

// WithRegisterTimeout is setting a timeout for registering with the controller.
// The timeout has to be > 0, otherwise it is the default of 3s.
func WithRegisterTimeout(timeout time.Duration) func(o *Opts) {
	return func(o *Opts) {
		o.RegisterTimeout = timeout
	}
}

// Events returns a channel which is used to publish important
// events to the plugin. For example that the server pushes a new config.
// or that the plugin needs to restart.
func (p *Plugin) Events() <-chan Event {
	out := make(chan Event)

	p.eventsChannels = append(p.eventsChannels, out)

	return out
}

// Log is returning the logger to log to the log formatter.
func (p *Plugin) Log() *log.Entry {
	return p.logger
}

func (p *Plugin) multiplexEvents() func() error {
	return func() error {
		for {
			select {
			case e := <-p.events:
				for _, c := range p.eventsChannels {
					// this pushes the events to a routine
					// to avoid the loop to block in execution
					p.Run(func() error {
						c <- e

						return nil
					})
				}
			case <-p.ctx.Done():
				return nil
			}
		}
	}
}

// newPlugin ...
func newPlugin(opts ...Opt) *Plugin {
	options := new(Opts)

	p := new(Plugin)
	// setting a default env for a plugin
	p.opts = options
	p.ready = make(chan struct{})
	p.events = make(chan Event)

	// configure plugin
	configure(p, opts...)

	// logging ...
	configureLogging(p)

	// starts the multiplexder
	p.Run(p.multiplexEvents())

	return p
}

// SubscribeInbox is subscribing to the inbox of messages.
// This if for plugins that want to consume the message that other
// plugins publish to Autobot.
func (p *Plugin) SubscribeInbox(funcs ...filters.FilterFunc) <-chan *pb.Event {
	sub := make(chan *pb.Event)

	p.Run(p.subInboxFunc(sub, funcs...))

	return sub
}

// SubscribeOutbox is subscribing to the outbox of messages.
// These are the message that ought to be published to an external service (e.g. Slack, MS Teams).
func (p *Plugin) SubscribeOutbox(funcs ...filters.FilterFunc) <-chan *pb.Event {
	sub := make(chan *pb.Event)

	p.Run(p.subOutboxFunc(sub, funcs...))

	return sub
}

// PublishInbox is publishing message to the inbox in the controller.
// The returned channel pushes all of the send message to the inbox in the controller.
func (p *Plugin) PublishInbox(funcs ...filters.FilterFunc) chan<- *pb.Event {
	pub := make(chan *pb.Event)

	p.Run(p.pubInboxFunc(pub, append(filters.DefaultInboxFilterOpts, funcs...)...))

	return pub
}

// PublishOutbox is publishing message to the outbox in the controller.
// The returned channel pushes all the send messages to the outbox in the controller.
func (p *Plugin) PublishOutbox(funcs ...filters.FilterFunc) chan<- *pb.Event {
	pub := make(chan *pb.Event)

	p.Run(p.pubOutboxFunc(pub, append(filters.DefaultOutboxFilterOpts, funcs...)...))

	return pub
}

// Wait is waiting for the underlying WaitGroup. All run go routines are
// hold here to before one exists with an error. Then the provided context is canceled.
func (p *Plugin) Wait() error {
	p.wg.Wait()

	if p.sc != nil {
		p.sc.Close()
	}

	if p.nc != nil {
		p.nc.Close()
	}

	return p.err
}

// ReplyWithFunc is a wrapper function to provide a function which may
// send replies to received message for this plugin.
func (p *Plugin) ReplyWithFunc(fn SubscribeFunc, funcs ...filters.FilterFunc) error {
	p.Run(func() error {
		// create publish channel ...
		pubReply := p.PublishOutbox()
		subMsg := p.SubscribeInbox()

		for {
			select {
			case e, ok := <-subMsg:
				if !ok {
					return nil
				}

				r, err := fn(e)
				if err != nil {
					return err
				}

				pubReply <- r
			case <-p.ctx.Done():
				return nil
			}
		}
	})

	return nil
}

// Run is a wrapper similar to errgroup to schedule functions
// as go routines in a waitGroup.
func (p *Plugin) Run(f func() error) {
	p.wg.Add(1)

	go func() {
		defer p.wg.Done()

		if err := f(); err != nil {
			p.errOnce.Do(func() {
				p.err = err
				if p.cancel != nil {
					p.cancel()
				}
			})
		}
	}()
}

func (p *Plugin) getConn() (stan.Conn, error) {
	p.Lock()
	defer p.Unlock()

	if p.sc != nil {
		return p.sc, nil
	}

	c, err := p.newConn()
	if err != nil {
		return nil, err
	}

	p.sc = c

	// start watcher ...
	p.Run(p.watch())

	if err := p.register(); err != nil {
		return nil, err
	}

	select {
	case <-time.After(p.opts.RegisterTimeout):
		return nil, ErrPluginRegister
	case <-p.ready:
		p.events <- ReadyEvent
	}

	return p.sc, nil
}

func (p *Plugin) newConn() (stan.Conn, error) {
	nc, err := nats.Connect(
		p.env.ClusterURL,
		nats.MaxReconnects(-1),
		nats.ReconnectBufSize(-1),
	)
	if err != nil {
		return nil, err
	}

	p.nc = nc

	sc, err := stan.Connect(
		p.env.ClusterID,
		p.env.Name,
		stan.NatsConn(nc),
		stan.SetConnectionLostHandler(p.lostHandler()),
	)
	if err != nil {
		return nil, err
	}

	// creates a new re-used response inbox for the plugin
	p.resp = p.nc.NewRespInbox()

	return sc, nil
}

func (p *Plugin) watch() func() error {
	return func() error {
		exit := make(chan error)
		resp := make(chan *pb.Event)

		sub, err := p.nc.Subscribe(p.resp, func(msg *nats.Msg) {
			r := new(pb.Event)
			if err := proto.Unmarshal(msg.Data, r); err != nil {
				// no nothing nows
				exit <- err
			}

			resp <- r
		})
		if err != nil {
			return err
		}

		defer sub.Unsubscribe()

		for {
			select {
			case err := <-exit:
				return err
			case <-p.ctx.Done():
				return nil
			case event := <-resp:
				if err := p.handleEvent(event); err != nil {
					return err
				}
			}
		}
	}
}

func (p *Plugin) handleEvent(action *pb.Event) error {
	// identify action ...
	switch a := action.GetEvent().(type) {
	case *pb.Event_Config:
		return p.handleConfig(a.Config)
	default:
	}

	return nil
}

func (p *Plugin) handleConfig(cfg *pb.Config) error {
	// reconfiguring plugin ...
	p.env = p.env.WithConfig(cfg)

	p.ready <- struct{}{}
	p.events <- RefreshConfigEvent

	return nil
}

func (p *Plugin) publishEvent(a *pb.Event) error {
	// try to marshal into []byte ...
	msg, err := proto.Marshal(a)
	if err != nil {
		return err
	}

	err = p.nc.PublishMsg(&nats.Msg{
		Subject: p.env.ClusterDiscovery,
		Reply:   p.resp,
		Data:    msg,
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *Plugin) register() error {
	plugin := &pb.Plugin{
		Name: p.env.Name,
	}

	// send the action
	if err := p.publishEvent(pb.NewRegister(plugin)); err != nil {
		return err
	}

	// set this to be registered
	return nil
}

func (p *Plugin) pubInboxFunc(pub <-chan *pb.Event, funcs ...filters.FilterFunc) func() error {
	return func() error {
		sc, err := p.getConn()
		if err != nil {
			return err
		}

		f := filters.New(funcs...)

		for {
			select {
			case e, ok := <-pub:
				if !ok {
					return nil
				}

				// filtering an event
				e, err := f.Filter(e)
				if err != nil || e == nil {
					return err
				}

				// try to marshal into []byte ...
				msg, err := proto.Marshal(e)
				if err != nil {
					return err
				}

				if err := sc.Publish(p.env.Inbox, msg); err != nil {
					return err
				}
			case <-p.ctx.Done():
				return nil
			}
		}
	}
}

func (p *Plugin) pubOutboxFunc(pub <-chan *pb.Event, funcs ...filters.FilterFunc) func() error {
	return func() error {
		sc, err := p.getConn()
		if err != nil {
			return err
		}

		f := filters.New(funcs...)

		for {
			select {
			case e, ok := <-pub:
				if !ok {
					return nil
				}

				if e == nil {
					continue
				}

				msg, err := proto.Marshal(e)
				if err != nil {
					return err
				}

				// filtering an event
				e, err = f.Filter(e)
				if err != nil {
					return err
				}

				// if this has not been filtered, or else
				if e == nil {
					continue
				}

				if err := sc.Publish(p.env.Outbox, msg); err != nil {
					return err
				}
			case <-p.ctx.Done():
				return nil
			}
		}
	}
}

func (p *Plugin) subInboxFunc(sub chan<- *pb.Event, funcs ...filters.FilterFunc) func() error {
	return func() error {
		sc, err := p.getConn()
		if err != nil {
			return err
		}

		f := filters.New(funcs...)

		s, err := sc.QueueSubscribe(p.env.Inbox, p.env.Name, func(m *stan.Msg) {
			event := new(pb.Event)
			if err := proto.Unmarshal(m.Data, event); err != nil {
				// no nothing now
				return
			}

			// filtering an event
			event, err := f.Filter(event)
			if err != nil {
				return
			}

			// if this has not been filtered, or else
			if event != nil {
				sub <- event
			}
		}, stan.DurableName(p.env.Name))
		if err != nil {
			return err
		}

		defer s.Close()

		<-p.ctx.Done()

		// close channel
		close(sub)

		return nil
	}
}

func (p *Plugin) subOutboxFunc(sub chan<- *pb.Event, funcs ...filters.FilterFunc) func() error {
	return func() error {
		sc, err := p.getConn()
		if err != nil {
			return err
		}

		f := filters.New(funcs...)

		s, err := sc.QueueSubscribe(p.env.Outbox, p.env.Name, func(m *stan.Msg) {
			event := new(pb.Event)
			if err := proto.Unmarshal(m.Data, event); err != nil {
				// no nothing now
				return
			}

			// filtering an event
			event, err := f.Filter(event)
			if err != nil {
				return
			}

			// if this has not been filtered, or else
			if event != nil {
				sub <- event
			}
		}, stan.DurableName(p.env.Name))
		if err != nil {
			return err
		}

		defer s.Close()

		<-p.ctx.Done()

		// close channel
		close(sub)

		return nil
	}
}

func (p *Plugin) lostHandler() func(_ stan.Conn, reason error) {
	return func(_ stan.Conn, reason error) {
		p.cancel()

		p.err = reason
	}
}

func configureLogging(p *Plugin) error {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	if p.env.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	p.logger = log.WithFields(
		log.Fields{
			"cluster_url": p.env.ClusterURL,
			"cluster_id":  p.env.ClusterID,
		},
	)

	if level, err := log.ParseLevel(p.env.LogLevel); err == nil {
		log.SetLevel(level)
	}

	return nil
}

func configure(p *Plugin, opts ...Opt) error {
	for _, o := range opts {
		o(p.opts)
	}

	if p.opts.RegisterTimeout == 0 {
		p.opts.RegisterTimeout = defaultRegisterTimeout
	}

	return nil
}
