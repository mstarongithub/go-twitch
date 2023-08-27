package events

import (
	"encoding/json"
	"strconv"
	"time"
)

// A Channel GuestStarSessionBegin Event.
// Sent when a guest star session begins
//
// Subscription name:
//
//	channel.guest_star_session.begin
//
// Required scope:
//
//	channel:read:guest_star
//
// or
//
//	channel:manage:guest_star
//
// # NOTE: Public Beta Event
type ChannelGuestStarSessionBeginEvent struct {
	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	SessionID string
	StartedAt time.Time
}

// Meta struct to convert from json full of strings to native types
type channelGuestStarSessionBeginEventMeta struct {
	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	SessionID string `json:"session_id"`
	StartedAt string `json:"started_at"`
}

func (c *ChannelGuestStarSessionBeginEvent) UnmarshalJSON(b []byte) error {
	var m channelGuestStarSessionBeginEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.BroadcasterID, err = strconv.Atoi(m.BroadcasterID)
	if err != nil {
		return err
	}
	c.BroadcasterName = m.BroadcasterName
	c.BroadcasterDisplayName = m.BroadcasterDisplayName

	c.SessionID = m.SessionID
	c.StartedAt, err = time.Parse(time.RFC3339Nano, m.StartedAt)
	if err != nil {
		return err
	}

	return nil
}
