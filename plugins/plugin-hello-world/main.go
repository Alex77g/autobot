package main

import (
	"log"

	"github.com/andersnormal/autobot/pkg/plugins"
	pb "github.com/andersnormal/autobot/proto"
)

func main() {
	// plugin ....
	plugin, err := plugins.New("hello-world")
	if err != nil {
		log.Fatalf("could not create plugin: %v", err)
	}

	// use the schedule function from the plugin
	if err := plugin.SubscribeMessage(msgFunc()); err != nil {
		log.Fatalf("could not create plugin: %v", err)
	}

	if err := plugin.Wait(); err != nil {
		log.Fatalf("stopped plugin: %v", err)
	}
}

func msgFunc() plugins.SubscribeFunc {
	return func(in *pb.Event) (*pb.Event, error) {
		if in.GetMessage() != nil {
			return &pb.Event{
				Event: &pb.Event_Reply{
					Reply: &pb.MessageEvent{
						Text:     "hello world",
						Channel:  in.GetMessage().GetChannel(),
						User:     in.GetMessage().GetUser(),
						Username: in.GetMessage().GetUsername(),
						Topic:    in.GetMessage().GetTopic(),
					},
				},
			}, nil
		}

		return nil, nil
	}
}
