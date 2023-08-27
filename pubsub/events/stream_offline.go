package events

import (
	"encoding/json"
	"strconv"
)

const STREAM_OFFLINE_EVENT = "stream.offline"

// A Stream Offline Event.
//
// Sent when the target channel goes offline.
//
// Subscription name:
//
//	stream.offline
//
// Required scope:
//
//	None
type StreamOfflineEvent struct {
	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string
}

// Meta struct to convert from json full of strings to native types
type streamOfflineEventMeta struct {
	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`
}

func (c *StreamOfflineEvent) UnmarshalJSON(b []byte) error {
	var m streamOfflineEventMeta
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

	return nil
}
