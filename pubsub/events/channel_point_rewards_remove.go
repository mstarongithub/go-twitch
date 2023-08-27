package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_POINTS_REWARDS_REMOVE_EVENT = "channel.channel_points_custom_reward.remove"

// A Channel PointRewardsRemove Event.
// Sent when a channel points reward is removed
//
// Subscription name:
//
//	channel.channel_points_custom_reward.remove
//
// Required scope:
//
//	channel:read:redemptions
//
// or
//
//	channel:manage:redemptions
type ChannelPointRewardsRemoveEvent struct {
	ID int

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Enabled                bool
	Paused                 bool
	InStock                bool
	Title                  string
	Cost                   int
	Prompt                 string
	UserInputRequired      bool
	ShouldSkipRequestQueue bool
	CooldownExpiresAt      *time.Time
	RedeemedCurrentStream  *int
	MaxPerStream           MaxPerStream
	MaxPerUserPerStream    MaxPerStream
	GlobalCooldown         struct {
		Enabled bool
		Seconds int
	}
	BackgroundColor string
	Image           Image
	DefaultImage    Image
}

// Meta struct to convert from json full of strings to native types
type channelPointRewardsRemoveEventMeta struct {
	ID string `json:"id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Enabled                bool         `json:"is_enabled"`
	Paused                 bool         `json:"is_paused"`
	InStock                bool         `json:"is_in_stock"`
	Title                  string       `json:"title"`
	Cost                   int          `json:"cost"`
	Prompt                 string       `json:"prompt"`
	UserInputRequired      bool         `json:"is_user_input_required"`
	ShouldSkipRequestQueue bool         `json:"should_redemptions_skip_request_queue"`
	CooldownExpiresAt      *time.Time   `json:"cooldown_expires_at:omitempty"`
	RedeemedCurrentStream  *int         `json:"redemptions_redeemed_current_stream:omitempty"`
	MaxPerStream           MaxPerStream `json:"max_per_stream"`
	MaxPerUserPerStream    MaxPerStream `json:"max_per_user_per_stream"`
	GlobalCooldown         struct {
		Enabled bool `json:"is_enabled"`
		Seconds int  `json:"seconds"`
	} `json:"global_cooldown"`
	BackgroundColor string `json:"background_color"`
	Image           Image  `json:"image"`
	DefaultImage    Image  `json:"default_image"`
}

func (c *ChannelPointRewardsRemoveEvent) UnmarshalJSON(b []byte) error {
	var m channelPointRewardsRemoveEventMeta
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

	c.Enabled = m.Enabled
	c.Paused = m.Paused
	c.InStock = m.InStock
	c.Title = m.Title
	c.Cost = m.Cost
	c.Prompt = m.Prompt
	c.UserInputRequired = m.UserInputRequired
	c.ShouldSkipRequestQueue = m.ShouldSkipRequestQueue
	c.CooldownExpiresAt = m.CooldownExpiresAt
	c.RedeemedCurrentStream = m.RedeemedCurrentStream
	c.MaxPerStream = m.MaxPerStream
	c.MaxPerUserPerStream = m.MaxPerUserPerStream
	c.GlobalCooldown = struct {
		Enabled bool
		Seconds int
	}(m.GlobalCooldown)
	c.BackgroundColor = m.BackgroundColor
	c.Image = m.Image
	c.DefaultImage = m.DefaultImage

	return nil
}
