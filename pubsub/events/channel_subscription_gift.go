package events

import (
	"encoding/json"
	"strconv"
)

// A Channel Subscription Gifted Event.
//
// Sent when a user gifts one or more subscriptions to the target channel.
//
// Subscription name:
//
//	channel.subscription.gift
//
// Required scope:
//
//	channel:read:subscriptions
type ChannelSubscriptionGiftedEvent struct {
	UserID          int
	UserName        string
	UserDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Total           int
	Tier            int
	CumulativeTotal *int
	Anonymous       bool
}

// Meta struct to convert from json full of strings to native types
type channelSubscriptionGiftedEventMeta struct {
	UserID          string `json:"user_id"`
	UserName        string `json:"user_login"`
	UserDisplayName string `json:"user_name"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Total           int    `json:"total"`
	Tier            string `json:"tier"`
	CumulativeTotal *int   `json:"cumulative_total:omitempty"`
	Anonymous       bool   `json:"is_anonymous"`
}

func (c *ChannelSubscriptionGiftedEvent) UnmarshalJSON(b []byte) error {
	var m channelSubscriptionGiftedEventMeta
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

	c.Total = m.Total
	c.Tier, err = strconv.Atoi(m.Tier)
	if err != nil {
		return err
	}
	c.CumulativeTotal = m.CumulativeTotal
	c.Anonymous = m.Anonymous

	return nil
}
