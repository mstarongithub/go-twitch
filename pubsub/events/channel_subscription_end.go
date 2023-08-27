package events

import (
	"encoding/json"
	"strconv"
)

const CHANNEL_SUBSCRIPTION_END_EVENT = "channel.subscription.end"

// A Channel Subscription End Event
//
// # Sent when a subscription to a target channel expires
//
// Subscription name
//
//	channel.subscription.end
//
// Required scope:
//
//	channel:read:subscriptions
type ChannelSubscriptionEndEvent struct {
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
type channelSubscriptionEndEventMeta struct {
	UserID          string `json:"user_id"`
	UserName        string `json:"user_login"`
	UserDisplayName string `json:"user_name"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Tier   string `json:"tier"`
	Gifted bool   `json:"is_gift"`
}

func (c *ChannelSubscriptionEndEvent) UnmarshalJSON(b []byte) error {
	var m channelSubscriptionEndEventMeta
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
