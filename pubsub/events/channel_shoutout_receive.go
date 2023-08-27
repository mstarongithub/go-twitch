package events

import (
	"encoding/json"
	"strconv"
	"time"
)

// A Channel ShoutoutReceive Event.
// Sent when the target channel receives a shoutout
//
// Subscription name:
//
//	channel.shoutout.receive
//
// Required scope:
//
//	moderator:read:shoutouts
//
// or
//
//	moderator:manage:shoutouts
//
// Note: The moderator id must match the user ID in the access token used.
// Note: Sent only if Twitch posts the Shoutout to the broadcaster’s activity feed
type ChannelShoutoutReceiveEvent struct {
	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	FromBroadcasterID          int
	FromBroadcasterName        string
	FromBroadcasterDisplayName string

	StartedAt   time.Time
	ViewerCount int
}

// Meta struct to convert from json full of strings to native types
type channelShoutoutReceiveEventMeta struct {
	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	FromBroadcasterID          string `json:"from_broadcaster_user_id"`
	FromBroadcasterName        string `json:"from_broadcaster_user_name"`
	FromBroadcasterDisplayName string `json:"from_broadcaster_user_login"`

	StartedAt   time.Time `json:"started_at"`
	ViewerCount int       `json:"viewer_count"`
}

func (c *ChannelShoutoutReceiveEvent) UnmarshalJSON(b []byte) error {
	var m channelShoutoutReceiveEventMeta
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

	c.FromBroadcasterID, err = strconv.Atoi(m.FromBroadcasterID)
	if err != nil {
		return err
	}
	c.FromBroadcasterName = m.FromBroadcasterName
	c.FromBroadcasterDisplayName = m.FromBroadcasterDisplayName

	c.StartedAt = m.StartedAt
	c.ViewerCount = m.ViewerCount

	return nil
}
