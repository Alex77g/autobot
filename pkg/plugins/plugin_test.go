package plugins

import (
	"context"
	"testing"
	"time"

	"github.com/andersnormal/autobot/pkg/plugins/runtime"
	pb "github.com/andersnormal/autobot/proto"

	"github.com/stretchr/testify/assert"
)

func TestInbox(t *testing.T) {
	env := &runtime.Environment{
		ClusterID:  runtime.DefaultClusterID,
		ClusterURL: runtime.DefaultClusterURL,
		Inbox:      runtime.DefaultClusterInbox,
		Outbox:     runtime.DefaultClusterOutbox,
		LogFormat:  runtime.DefaultLogFormat,
		LogLevel:   runtime.DefaultLogLevel,
	}

	env.Name = "skrimish"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	withTestAutobot(ctx, t, env, func(t *testing.T, cancel context.CancelFunc) {
		assert := assert.New(t)
		defer cancel()

		// create test plugin ....
		plugin, _ := WithContext(ctx, env)

		// create channels...
		subMsg := plugin.SubscribeInbox()

		err := plugin.PublishInbox(&pb.Message{Text: "foo"})
		assert.NoError(err)

		select {
		case e, ok := <-subMsg:
			assert.True(ok)

			assert.IsType(e, &pb.Message{})
			assert.Equal(e.(*pb.Message).GetText(), "foo")
		case <-time.After(5 * time.Second):
			assert.FailNow("timed out waiting for message to arrive at the outbox")
		}
	})
}

func TestOutbox(t *testing.T) {
	env := &runtime.Environment{
		ClusterID:  runtime.DefaultClusterID,
		ClusterURL: runtime.DefaultClusterURL,
		Inbox:      runtime.DefaultClusterInbox,
		Outbox:     runtime.DefaultClusterOutbox,
		LogFormat:  runtime.DefaultLogFormat,
		LogLevel:   runtime.DefaultLogLevel,
	}

	env.Name = "skrimish"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	withTestAutobot(ctx, t, env, func(t *testing.T, cancel context.CancelFunc) {
		assert := assert.New(t)
		defer cancel()

		// create test plugin ....
		plugin, _ := WithContext(ctx, env)

		// create channels...
		subMsg := plugin.SubscribeOutbox()

		err := plugin.PublishOutbox(&pb.Message{Text: "foo"})
		assert.NoError(err)

		select {
		case e, ok := <-subMsg:
			assert.True(ok)

			assert.IsType(e, &pb.Message{})
			assert.Equal(e.(*pb.Message).GetText(), "foo")
		case <-time.After(5 * time.Second):
			assert.FailNow("timed out waiting for message to arrive at the outbox")
		}
	})
}
