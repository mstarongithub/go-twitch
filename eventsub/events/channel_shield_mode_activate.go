package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_SHIELD_MODE_BEGIN_EVENT = "channel.shield_mode.begin"

// A Channel ShieldModeBegin Event.
// Sent when the shield mode is activated
//
// Subscription name:
//
//	channel.shield_mode.begin
//
// Required scope:
//
//	moderator:read:shield_mode
//
// or
//
//	moderator:manage:shield_mode
//
// Note: The moderator id must match the user ID in the access token used
type ChannelShieldModeBeginEvent struct {
	ModID          int
	ModName        string
	ModDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	StartedAt time.Time
}

// Meta struct to convert from json full of strings to native types
type channelShieldModeBeginEventMeta struct {
	ModID          string `json:"moderator_user_id"`
	ModName        string `json:"moderator_user_name"`
	ModDisplayName string `json:"moderator_user_login"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	StartedAt time.Time `json:"started_at"`
}

func (c *ChannelShieldModeBeginEvent) UnmarshalJSON(b []byte) error {
	var m channelShieldModeBeginEventMeta
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

	c.StartedAt = m.StartedAt

	return nil
}
