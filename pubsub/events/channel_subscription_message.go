package events

import (
	"encoding/json"
	"strconv"
)

const CHANNEL_SUBSCRIPTION_MESSAGE_EVENT = "channel.subscription.message"

// A Channel Subscription Message Event.
//
// Sent when a user sends a resubscription chat message to the target channel.
//
// Subscription name:
//
//	channel.subscription.message
//
// Required scope:
//
//	channel:read:subscriptions
type ChannelSubscriptionMessageEvent struct {
	UserID          int
	UserName        string
	UserDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Tier             int
	Message          EventMessage
	CumulativeMonths int
	Streak           *int
	Duration         int
}

// Meta struct to convert from json full of strings to native types
type channelSubscriptionMessageEventMeta struct {
	UserID          string `json:"user_id"`
	UserName        string `json:"user_login"`
	UserDisplayName string `json:"user_name"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Tier             string       `json:"tier"`
	Message          EventMessage `json:"message"`
	CumulativeMonths int          `json:"cumulative_months"`
	Streak           *int         `json:"streak_months"`
	Duration         int          `json:"duration_months"`
}

func (c *ChannelSubscriptionMessageEvent) UnmarshalJSON(b []byte) error {
	var m channelSubscriptionMessageEventMeta
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
	c.Message = m.Message
	c.CumulativeMonths = m.CumulativeMonths
	c.Streak = m.Streak
	c.Duration = m.Duration

	return nil
}
