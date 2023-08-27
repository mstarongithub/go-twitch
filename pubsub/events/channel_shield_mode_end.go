package events

import (
	"encoding/json"
	"strconv"
	"time"
)

// A Channel ShieldModeEnd Event.
// Sent when the shield mode ends
//
// Subscription name:
//
//	channel.shield_mode.end
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
type ChannelShieldModeEndEvent struct {
	ModID          int
	ModName        string
	ModDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	EndedAt time.Time
}

// Meta struct to convert from json full of strings to native types
type channelShieldModeEndEventMeta struct {
	ModID          string `json:"moderator_user_id"`
	ModName        string `json:"moderator_user_name"`
	ModDisplayName string `json:"moderator_user_login"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	EndedAt time.Time `json:"ended_at"`
}

func (c *ChannelShieldModeEndEvent) UnmarshalJSON(b []byte) error {
	var m channelShieldModeEndEventMeta
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

	c.EndedAt = m.EndedAt

	return nil
}
