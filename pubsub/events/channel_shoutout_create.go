package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_SHOUTOUT_CREATE_EVENT = "channel.shoutout.create"

// A Channel ShoutoutCreate Event.
// Sent when the target channel sends a shoutout
//
// Subscription name:
//
//	channel.shoutout.create
//
// Required scope:
//
//	moderator:read:shoutouts
//
// or
//
//	moderator:manage:shoutouts
//
// Note: The moderator id must match the user ID in the access token used
type ChannelShoutoutCreateEvent struct {
	ModID          int
	ModName        string
	ModDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	ToBroadcasterID          int
	ToBroadcasterName        string
	ToBroadcasterDisplayName string

	StartedAt            time.Time
	ViewerCount          int
	CooldownEndsAt       time.Time
	TargetCooldownEndsAt time.Time
}

// Meta struct to convert from json full of strings to native types
type channelShoutoutCreateEventMeta struct {
	ModID          string `json:"moderator_user_id"`
	ModName        string `json:"moderator_user_name"`
	ModDisplayName string `json:"moderator_user_login"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	ToBroadcasterID          string `json:"to_broadcaster_user_id"`
	ToBroadcasterName        string `json:"to_broadcaster_user_name"`
	ToBroadcasterDisplayName string `json:"to_broadcaster_user_login"`

	StartedAt            time.Time `json:"started_at"`
	ViewerCount          int       `json:"viewer_count"`
	CooldownEndsAt       time.Time `json:"cooldown_ends_at"`
	TargetCooldownEndsAt time.Time `json:"target_cooldown_ends_at"`
}

func (c *ChannelShoutoutCreateEvent) UnmarshalJSON(b []byte) error {
	var m channelShoutoutCreateEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.ModID, err = strconv.Atoi(m.ModID)
	if err != nil {
		return err
	}
	c.ModName = m.ModName
	c.ModDisplayName = m.ModDisplayName

	c.BroadcasterID, err = strconv.Atoi(m.BroadcasterID)
	if err != nil {
		return err
	}
	c.BroadcasterName = m.BroadcasterName
	c.BroadcasterDisplayName = m.BroadcasterDisplayName

	c.ToBroadcasterID, err = strconv.Atoi(m.ToBroadcasterID)
	if err != nil {
		return err
	}
	c.ToBroadcasterName = m.ToBroadcasterName
	c.ToBroadcasterDisplayName = m.ToBroadcasterDisplayName

	c.StartedAt = m.StartedAt
	c.ViewerCount = m.ViewerCount
	c.CooldownEndsAt = m.CooldownEndsAt
	c.TargetCooldownEndsAt = m.TargetCooldownEndsAt

	return nil
}
