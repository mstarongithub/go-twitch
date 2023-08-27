package events

import (
	"encoding/json"
	"strconv"
	"time"
)

// A Channel Ban Event.
// Sent when a user gets banned
//
// Subscription name:
//
//	channel.ban
//
// Required scope:
//
//	channel:moderate
type ChannelBanEvent struct {
	UserID          int
	UserName        string
	UserDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	ModID          int
	ModName        string
	ModDisplayName string

	Reason    string
	BannedAt  time.Time
	EndsAt    time.Time
	Permanent bool
}

// Meta struct to convert from json full of strings to native types
type channelBanEventMeta struct {
	UserID          string `json:"user_id"`
	UserName        string `json:"user_login"`
	UserDisplayName string `json:"user_name"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	ModID          string `json:"moderator_user_id"`
	ModName        string `json:"moderator_user_login"`
	ModDisplayName string `json:"moderator_user_name"`

	Reason    string    `json:"reason"`
	BannedAt  time.Time `json:"banned_at"`
	EndsAt    time.Time `json:"ends_at"`
	Permanent bool      `json:"is_permanent"`
}

func (c *ChannelBanEvent) UnmarshalJSON(b []byte) error {
	var m channelBanEventMeta
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

	c.ModID, err = strconv.Atoi(m.ModID)
	if err != nil {
		return err
	}
	c.ModName = m.ModName
	c.ModDisplayName = m.ModDisplayName

	c.Reason = m.Reason
	c.BannedAt = m.BannedAt
	c.EndsAt = m.EndsAt
	c.Permanent = m.Permanent

	return nil
}
