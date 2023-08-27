package events

import (
	"encoding/json"
	"strconv"
)

const CHANNEL_SUBSCRIBE_EVENT = "channel.subscribe"

// A Channel Subscription Event
//
// Sent when a a target channel receives a subscription
// Does not include resubscribes
//
// Subscription name:
//
//	channel.subscribe
//
// Required scope:
//
//	channel:read:subscriptions
type ChannelSubscribeEvent struct {
	UserID          int
	UserName        string
	UserDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Tier   int
	Gifted bool
}

// Meta struct to convert from json full of strings to native types
type channelSubscribeEventMeta struct {
	UserID          string `json:"user_id"`
	UserName        string `json:"user_login"`
	UserDisplayName string `json:"user_name"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Tier   string `json:"tier"`
	Gifted bool   `json:"is_gift"`
}

func (c *ChannelSubscribeEvent) UnmarshalJSON(b []byte) error {
	var m channelSubscribeEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.UserID, err = strconv.Atoi(m.UserID)
	if err != nil {
		return err
	}
	c.UserName = m.UserName
	c.UserDisplayName = m.UserDisplayName

	c.BroadcasterID, err = strconv.Atoi(m.BroadcasterID)
	if err != nil {
		return err
	}
	c.BroadcasterName = m.BroadcasterName
	c.BroadcasterDisplayName = m.BroadcasterDisplayName

	c.Tier, err = strconv.Atoi(m.Tier)
	if err != nil {
		return err
	}
	c.Gifted = m.Gifted

	return nil
}
