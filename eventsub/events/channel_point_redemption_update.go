package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_POINTS_REDEMPTION_UPDATE_EVENT = "channel.channel_points_custom_reward_redemption.update"

// A Channel PointsRedemptionUpdate Event.
// Sent when a redeemed channel points reward is updated
//
// Subscription name:
//
//	channel.channel_points_custom_reward_redemption.update
//
// Required scope:
//
//	channel:read:redemptions
//
// or
//
//	channel:manage:redemptions
type ChannelPointsRedemptionUpdateEvent struct {
	ID string

	UserID          int
	UserName        string
	UserDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Input  string
	Status string // Either fulfilled or cancelled
	Reward struct {
		ID     string
		Title  string
		Cost   int
		Prompt string
	}
	RedeemedAt time.Time
}

// Meta struct to convert from json full of strings to native types
type channelPointsRedemptionUpdateEventMeta struct {
	ID                     string `json:"id"`
	UserID                 string `json:"user_id"`
	UserName               string `json:"user_login"`
	UserDisplayName        string `json:"user_name"`
	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`
	Input                  string `json:"user_input"`
	Status                 string `json:"status"`
	Reward                 struct {
		ID     string `json:"id"`
		Title  string `json:"title"`
		Cost   int    `json:"cost"`
		Prompt string `json:"prompt"`
	} `json:"reward"`
	RedeemedAt time.Time `json:"redeemed_at"`
}

func (c *ChannelPointsRedemptionUpdateEvent) UnmarshalJSON(b []byte) error {
	var m channelPointsRedemptionUpdateEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.ID = m.ID

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

	c.Input = m.Input
	c.Status = m.Status
	c.Reward = struct {
		ID     string
		Title  string
		Cost   int
		Prompt string
	}(m.Reward)
	c.RedeemedAt = m.RedeemedAt

	return nil
}
