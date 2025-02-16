package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_PREDICTION_BEGIN_EVENT = "channel.prediction.begin"

// A Channel PredictionBegin Event.
// Sent when a prediction begins
//
// Subscription name:
//
//	channel.prediction.begin
//
// Required scope:
//
//	channel:read:predictions
//
// or
//
//	channel:manage:predictions
type ChannelPredictionBeginEvent struct {
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
type channelPredictionBeginEventMeta struct {
	ID string `json:"id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Title     string              `json:"title"`
	Outcomes  []PredictionOutcome `json:"outcomes"`
	StartedAt time.Time           `json:"started_at"`
	EndsAt    time.Time           `json:"ends_at"`
}

func (c *ChannelPredictionBeginEvent) UnmarshalJSON(b []byte) error {
	var m channelPredictionBeginEventMeta
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
	if err != nil {
		return err
	}
	c.EndsAt = m.EndsAt
	if err != nil {
		return err
	}

	return nil
}
