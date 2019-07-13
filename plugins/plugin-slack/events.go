package main

import (
	"context"
	"strconv"

	pb "github.com/andersnormal/autobot/proto"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/nlopes/slack"
)

// FromUserIDWithContext ...
func FromUserIDWithContext(ctx context.Context, api *slack.Client, userID string) (*pb.User, error) {
	u := new(pb.User)

	// get user
	user, err := api.GetUserInfoContext(ctx, userID)
	if err != nil {
		return nil, err
	}

	u.Id = user.ID
	u.Name = user.Name

	return u, nil
}

// FromChannelWithContext ...
func FromChannelIDWithContext(ctx context.Context, api *slack.Client, channelID string) (*pb.Channel, error) {
	c := new(pb.Channel)

	// get channel
	channel, err := api.GetChannelInfoContext(ctx, channelID)
	if err != nil {
		return nil, err
	}

	c.Id = channel.ID
	c.Name = channel.Name

	return c, nil
}

// FromChannelID ...
func FromChannelID(channelID string) *pb.Channel {
	return &pb.Channel{Id: channelID}
}

// FromMsgWithContext ...
func FromMsgWithContext(ctx context.Context, api *slack.Client, msg *slack.MessageEvent) (*pb.Event, error) {
	m := new(pb.Message)

	// basic ...
	m.Type = msg.Type

	// resolve user
	user, err := FromUserIDWithContext(ctx, api, msg.User)
	if err != nil {
		return nil, err
	}

	m.From = user

	// get list of direct message channels
	im, err := api.GetIMChannelsContext(ctx)
	if err != nil {
		return nil, err
	}

	var isPrivate bool
	for _, dm := range im {
		isPrivate = dm.ID == msg.Channel

		if isPrivate {
			m.Channel = FromChannelID(dm.ID)

			break
		}
	}

	if !isPrivate {
		// resolve channel ...
		channel, err := FromChannelIDWithContext(ctx, api, msg.Channel)
		if err != nil {
			return nil, err
		}

		m.Channel = channel
	}

	m.TextFormat = pb.Message_PLAIN_TEXT
	m.Text = msg.Text

	// timestamp ...
	ts, err := FromTimestamp(msg.Timestamp)
	if err != nil {
		return nil, err
	}
	m.Timestamp = ts

	// standard event ...
	e := &pb.Event{Event: &pb.Event_Message{
		Message: m,
	}}

	// if it is private
	if isPrivate {
		e.Event = &pb.Event_Private{
			Private: m,
		}
	}

	return e, nil
}

// FromTimestamp ...
func FromTimestamp(ts string) (*timestamp.Timestamp, error) {
	f, err := strconv.ParseFloat(ts, 64)
	if err != nil {
		return nil, err
	}

	secs := int64(f)
	nsecs := int32((f - float64(secs)) * 1e9)

	return &timestamp.Timestamp{
		Seconds: secs,
		Nanos:   nsecs,
	}, nil
}

// FromReplyEvent ...
func FromReplyEvent(rtm *slack.RTM, e *pb.Message) *slack.OutgoingMessage {
	return rtm.NewOutgoingMessage(e.GetText(), e.GetChannel().GetId())
}