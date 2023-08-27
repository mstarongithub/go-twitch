package events

import (
	"encoding/json"
	"strconv"
)

// A Channel ModPromotion Event.
// Sent when a user gets promoted to a moderator
//
// Subscription name:
//
//	channel.moderator.add
//
// Required scope:
//
//	moderation:read
type ChannelModPromotionEvent struct {
	UserID          int
	UserName        string
	UserDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string
}

// Meta struct to convert from json full of strings to native types
type channelModPromotionEventMeta struct {
	UserID          string `json:"user_id"`
	UserName        string `json:"user_login"`
	UserDisplayName string `json:"user_name"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`
}

func (c *ChannelModPromotionEvent) UnmarshalJSON(b []byte) error {
	var m channelModPromotionEventMeta
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

	return nil
}
