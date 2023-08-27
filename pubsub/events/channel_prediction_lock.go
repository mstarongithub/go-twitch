package events

import (
	"encoding/json"
	"strconv"
	"time"
)

// A Channel PredictionLock Event.
// Sent when a prediction is locked
//
// Subscription name:
//
//	channel.prediction.lock
//
// Required scope:
//
//	channel:read:predictions
//
// or
//
//	channel:manage:predictions
type ChannelPredictionLockEvent struct {
	ID string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Title     string
	Outcomes  []PredictionOutcome
	StartedAt time.Time
	EndsAt    time.Time
}

// Meta struct to convert from json full of strings to native types
type channelPredictionLockEventMeta struct {
	ID string `json:"id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Title     string              `json:"title"`
	Outcomes  []PredictionOutcome `json:"outcomes"`
	StartedAt time.Time           `json:"started_at"`
	EndsAt    time.Time           `json:"ends_at"`
}

func (c *ChannelPredictionLockEvent) UnmarshalJSON(b []byte) error {
	var m channelPredictionLockEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.ID = m.ID

	c.BroadcasterID, err = strconv.Atoi(m.BroadcasterID)
	if err != nil {
		return err
	}
	c.BroadcasterName = m.BroadcasterName
	c.BroadcasterDisplayName = m.BroadcasterDisplayName

	c.Title = m.Title
	c.Outcomes = m.Outcomes
	c.StartedAt = m.StartedAt
	c.EndsAt = m.EndsAt

	return nil
}
