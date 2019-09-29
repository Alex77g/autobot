package runtime

import (
	"errors"
)

var (
	// ErrNoRuntimeFunc signals that no functions have been configured to be run
	ErrNoRuntimeFunc = errors.New("runtime: not runtime functions have been configured")
)

const (
	// DefaultClusterID ...
	DefaultClusterID = "autobot"
	// DefaultClusterURL ...
	DefaultClusterURL = "nats://localhost:4222"
	// DefaultClusterInbox ...
	DefaultClusterInbox = "autobot.inbox"
	// DefaultClusterOutbox ...
	DefaultClusterOutbox = "autobot.outbox"
	// DefaultLogFormat ...
	DefaultLogFormat = "text"
	// DefaultLogLevel ...
	DefaultLogLevel = "warn"
)

var initializers []func()
var env *Environment

func init() {
	env = &Environment{
		ClusterID:  DefaultClusterID,
		ClusterURL: DefaultClusterURL,
		Inbox:      DefaultClusterInbox,
		Outbox:     DefaultClusterOutbox,
		LogFormat:  DefaultLogFormat,
		LogLevel:   DefaultLogLevel,
	}
}

// Env returns the current configured runtime environment.
func Env() *Environment {
	return env
}

// OnInitialize sets the passed functions to be run when runtime
// is called for initialization.
func OnInitialize(y ...func()) {
	initializers = append(initializers, y...)
}

func runInitializers() {
	for _, fn := range initializers {
		fn()
	}
}

// Runtime is a plugin runtime that executes runtime functions
type Runtime struct {
	Run  func(*Environment)
	RunE func(*Environment) error
}

// NewRuntime is returning a new Runtime
func NewRuntim() *Runtime {
	return &Runtime{}
}

// Environment describes a runtime environment for a plugin.
// This contains information about the used NATS cluster,
// the cluster id and the topic for plugin discovery.
type Environment struct {
	ClusterID  string `mapstructure:"cluster_id"`
	ClusterURL string `mapstructure:"cluster_url"`
	Debug      bool
	Inbox      string
	LogFormat  string `mapstructure:"log_format"`
	LogLevel   string `mapstructure:"log_level"`
	Name       string
	Outbox     string
	Verbose    bool
}

func (r *Runtime) hasRuntimeFuncs() bool {
	return r.Run != nil || r.RunE != nil
}

// Execute is running the configured runtime functions.
// It checks if there are functions configured and executes them
// with the configured environment. The initializer functions are run before
// the execution.
func (r *Runtime) Execute() error {
	if !r.hasRuntimeFuncs() {
		return ErrNoRuntimeFunc
	}

	runInitializers()

	if r.Run != nil {
		r.Run(env)

		return nil
	}

	if r.RunE != nil {
		if err := r.RunE(env); err != nil {
			return err
		}
	}

	return nil
}
