package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_HYPE_TRAIN_PROGRESS_EVENT = "channel.hype_train.progress"

// A Channel HypeTrainProgress Event.
// Sent when a HypeTrain progresses.
// May arrive out of order
//
// Subscription name:
//
//	channel.hype_train.progress
//
// Required scope:
//
//	channel:read:hype_train
type ChannelHypeTrainProgressEvent struct {
	ID string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Total            int
	Progress         int
	Goal             int
	TopContributions []HypeTrainContribution
	LastContribution HypeTrainContribution
	Level            int
	StartedAt        time.Time
	ExpiresAt        time.Time
}

// Meta struct to convert from json full of strings to native types
type channelHypeTrainProgressEventMeta struct {
	Id string `json:"id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Total            int                     `json:"total"`
	Progress         int                     `json:"progress"`
	Goal             int                     `json:"goal"`
	TopContributions []HypeTrainContribution `json:"top_contributions"`
	LastContribution HypeTrainContribution   `json:"last_contribution"`
	Level            int                     `json:"level"`
	StartedAt        time.Time               `json:"started_at"`
	ExpiresAt        time.Time               `json:"expires_at"`
}

func (c *ChannelHypeTrainProgressEvent) UnmarshalJSON(b []byte) error {
	var m channelHypeTrainProgressEventMeta
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

	c.Total = m.Total
	c.Progress = m.Progress
	c.Goal = m.Goal
	c.TopContributions = m.TopContributions
	c.LastContribution = m.LastContribution
	c.Level = m.Level
	c.StartedAt = m.StartedAt
	c.ExpiresAt = m.ExpiresAt

	return nil
}
