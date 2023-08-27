package events

import (
	"encoding/json"
	"strconv"
	"time"
)

// A Stream Online Event.
//
// Sent when the target channel goes online.
//
// Subscription name:
//
//	stream.online
//
// Required scope:
//
//	None
type StreamOnlineEvent struct {
	ID int

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Type      string
	StartedAt time.Time
}

// Meta struct to convert from json full of strings to native types
type streamOnlineEventMeta struct {
	ID                     string
	BroadcasterID          string    `json:"broadcaster_user_id"`
	BroadcasterName        string    `json:"broadcaster_user_login"`
	BroadcasterDisplayName string    `json:"broadcaster_user_name"`
	Type                   string    `json:"type"`
	StartedAt              time.Time `json:"started_at"`
}

func (c *StreamOnlineEvent) UnmarshalJSON(b []byte) error {
	var m streamOnlineEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.ID, err = strconv.Atoi(m.ID)
	if err != nil {
		return err
	}
	c.BroadcasterID, err = strconv.Atoi(m.BroadcasterID)
	if err != nil {
		return err
	}
	c.BroadcasterName = m.BroadcasterName
	c.BroadcasterDisplayName = m.BroadcasterDisplayName
	c.Type = m.Type
	c.StartedAt = m.StartedAt

	return nil
}
