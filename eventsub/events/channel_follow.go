package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_FOLLOW_EVENT = "channel.follow"

// A Channel Follow Event.
// Sent when a target channel receives a follow
//
// Subscription name:
//
//	channel.follow
//
// Required scope:
//
//	moderator:read:followers
type ChannelFollowEvent struct {
	UserID          int
	UserName        string
	UserDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	FollowedAt time.Time
}

// Meta struct to convert from json full of strings to native types
type channelFollowEventMeta struct {
	UserID          string `json:"user_id"`
	UserName        string `json:"user_login"`
	UserDisplayName string `json:"user_name"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	FollowedAt time.Time `json:"followed_at"`
}

func (c *ChannelFollowEvent) UnmarshalJSON(b []byte) error {
	var m channelFollowEventMeta
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

	c.FollowedAt = m.FollowedAt

	return nil
}
