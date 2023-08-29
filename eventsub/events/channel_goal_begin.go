package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_GOAL_BEGIN_EVENT = "channel.goal.begin"

// A Channel GoalBegin Event.
// Sent when a Goal begins.
// Not guaranteed to arrive before a ChannelGoalProgressEvent event
//
// Subscription name:
//
//	channel.goal.begin
//
// Required scope:
//
//	channel:read:goals
type ChannelGoalBeginEvent struct {
	ID string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Type          string
	Description   string
	CurrentAmount int
	TargetAmount  int
	StartedAt     time.Time
}

// Meta struct to convert from json full of strings to native types
type channelGoalBeginEventMeta struct {
	Id string `json:"id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Type          string    `json:"type"`
	Description   string    `json:"description"`
	CurrentAmount int       `json:"current_amount"`
	TargetAmount  int       `json:"target_amount"`
	StartedAt     time.Time `json:"started_at"`
}

func (c *ChannelGoalBeginEvent) UnmarshalJSON(b []byte) error {
	var m channelGoalBeginEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.ID = m.Id

	c.BroadcasterID, err = strconv.Atoi(m.BroadcasterID)
	if err != nil {
		return err
	}
	c.BroadcasterName = m.BroadcasterName
	c.BroadcasterDisplayName = m.BroadcasterDisplayName

	c.Type = m.Type
	c.Description = m.Description
	c.CurrentAmount = m.CurrentAmount
	c.TargetAmount = m.TargetAmount
	c.StartedAt = m.StartedAt

	return nil
}
