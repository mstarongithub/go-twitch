package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_PREDICTION_PROGRESS_EVENT = "channel.prediction.progress"

// A Channel PredictionProgress Event.
// Sent when a user participates in a prediction
//
// Subscription name:
//
//	channel.prediction.progress
//
// Required scope:
//
//	channel:read:predictions
//
// or
//
//	channel:manage:predictions
type ChannelPredictionProgressEvent struct {
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
type channelPredictionProgressEventMeta struct {
	ID string `json:"id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Title     string              `json:"title"`
	Outcomes  []PredictionOutcome `json:"outcomes"`
	StartedAt time.Time           `json:"started_at"`
	EndsAt    time.Time           `json:"ends_at"`
}

func (c *ChannelPredictionProgressEvent) UnmarshalJSON(b []byte) error {
	var m channelPredictionProgressEventMeta
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
