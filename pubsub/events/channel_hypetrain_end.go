package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_HYPE_TRAIN_END_EVENT = "channel.hype_train.end"

// A Channel HypeTrainEnd Event.
// Sent when a HypeTrain ends
//
// Subscription name:
//
//	channel.hype_train.end
//
// Required scope:
//
//	channel:read:hype_train
type ChannelHypeTrainEndEvent struct {
	ID                     string
	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string
	Total                  int
	TopContributions       []HypeTrainContribution
	LastContribution       HypeTrainContribution
	Level                  int
	StartedAt              time.Time
	EndedAt                time.Time
	CoolDownEndsAt         time.Time
}

// Meta struct to convert from json full of strings to native types
type channelHypeTrainEndEventMeta struct {
	Id string `json:"id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Level            int                     `json:"level"`
	Total            int                     `json:"total"`
	TopContributions []HypeTrainContribution `json:"top_contributions"`
	StartedAt        time.Time               `json:"started_at"`
	EndedAt          time.Time               `json:"ended_at"`
	CoolDownEndsAt   time.Time               `json:"cooldown_ends_at"`
}

func (c *ChannelHypeTrainEndEvent) UnmarshalJSON(b []byte) error {
	var m channelHypeTrainEndEventMeta
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
	c.TopContributions = m.TopContributions
	c.Level = m.Level
	c.StartedAt = m.StartedAt
	c.EndedAt = m.EndedAt
	c.CoolDownEndsAt = m.CoolDownEndsAt

	return nil
}
