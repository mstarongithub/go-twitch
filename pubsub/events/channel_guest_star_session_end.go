package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_GUEST_STAR_SESSION_END_EVENT = "channel.guest_star_session.end"

// A Channel GuestStarSessionEnd Event.
// Sent when a guest star session ends
//
// Subscription name:
//
//	channel.guest_star_session.end
//
// Required scope:
//
//	channel:read:guest_star
//
// or
//
//	channel:manage:guest_star
//
// or
//
//	moderator:read:guest_star
//
// or
//
//	moderator:manage:guest_star
//
// # NOTE: Public Beta Event
type ChannelGuestStarSessionEndEvent struct {
	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	SessionID string
	StartedAt time.Time
	EndedAt   time.Time
}

// Meta struct to convert from json full of strings to native types
type channelGuestStarSessionEndEventMeta struct {
	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	SessionID string `json:"session_id"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}

func (c *ChannelGuestStarSessionEndEvent) UnmarshalJSON(b []byte) error {
	var m channelGuestStarSessionEndEventMeta
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
	c.EndedAt, err = time.Parse(time.RFC3339Nano, m.EndedAt)
	if err != nil {
		return err
	}

	return nil
}
