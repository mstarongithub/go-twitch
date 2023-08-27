package events

import (
	"encoding/json"
	"strconv"
	"time"
)

// A Channel GoalEnd Event.
// Sent when a Goal has ended.
//
// Subscription name:
//
//	channel.goal.end
//
// Required scope:
//
//	channel:read:goals
type ChannelGoalEndEvent struct {
	ID string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Type          string
	Description   string
	Achieved      bool
	CurrentAmount int
	TargetAmount  int
	StartedAt     time.Time
	EndedAt       time.Time
}

// Meta struct to convert from json full of strings to native types
type channelGoalEndEventMeta struct {
	Id string `json:"id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Type          string    `json:"type"`
	Description   string    `json:"description"`
	Achieved      bool      `json:"is_achieved"`
	CurrentAmount int       `json:"current_amount"`
	TargetAmount  int       `json:"target_amount"`
	StartedAt     time.Time `json:"started_at"`
	EndedAt       time.Time `json:"ended_at"`
}

func (c *ChannelGoalEndEvent) UnmarshalJSON(b []byte) error {
	var m channelGoalEndEventMeta
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
	c.Achieved = m.Achieved
	c.CurrentAmount = m.CurrentAmount
	c.TargetAmount = m.TargetAmount
	c.StartedAt = m.StartedAt
	c.EndedAt = m.EndedAt

	return nil
}
