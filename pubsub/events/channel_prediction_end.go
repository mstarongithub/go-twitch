package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_PREDICTION_END_EVENT = "channel.prediction.end"

// A Channel PredictionEnd Event.
// Sent when a prediction ends
//
// Subscription name:
//
//	channel.prediction.end
//
// Required scope:
//
//	channel:read:predictions
//
// or
//
//	channel:manage:predictions
type ChannelPredictionEndEvent struct {
	ID string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Title            string
	WinningOutcomeID string
	Outcomes         []PredictionOutcome
	Status           string
	StartedAt        time.Time
	EndedAt          time.Time
}

// Meta struct to convert from json full of strings to native types
type channelPredictionEndEventMeta struct {
	ID string `json:"id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Title            string              `json:"title"`
	WinningOutcomeID string              `json:"winning_outcome_id"`
	Outcomes         []PredictionOutcome `json:"outcomes"`
	Status           string              `json:"status"`
	StartedAt        time.Time           `json:"started_at"`
	EndedAt          time.Time           `json:"ended_at"`
}

func (c *ChannelPredictionEndEvent) UnmarshalJSON(b []byte) error {
	var m channelPredictionEndEventMeta
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
	c.WinningOutcomeID = m.WinningOutcomeID
	c.Outcomes = m.Outcomes
	c.Status = m.Status
	c.StartedAt = m.StartedAt
	c.EndedAt = m.EndedAt

	return nil
}
