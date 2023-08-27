package events

import (
	"encoding/json"
	"strconv"
	"time"
)

// A Channel PollBegin Event.
// Sent when a poll is started on the target channel
//
// Subscription name:
//
//	channel.poll.begin
//
// Required scope:
//
//	channel:read:polls
//
// or
//
//	channel:manage:polls
type ChannelPollBeginEvent struct {
	ID string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Title               string
	Choices             []PollChoice
	BitsVoting          VotingCondition
	ChannelPointsVoting VotingCondition
	StartedAt           time.Time
	EndsAt              time.Time
}

// Meta struct to convert from json full of strings to native types
type channelPollBeginEventMeta struct {
	ID string `json:"id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Title               string          `json:"title"`
	Choices             []PollChoice    `json:"choices"`
	BitsVoting          VotingCondition `json:"bits_voting"`
	ChannelPointsVoting VotingCondition `json:"channel_points_voting"`
	StartedAt           time.Time       `json:"started_at"`
	EndsAt              time.Time       `json:"ends_at"`
}

type VotingCondition struct {
	Enabled       bool `json:"is_enabled"`
	AmountPerVote int  `json:"amount_per_vote"`
}

func (c *ChannelPollBeginEvent) UnmarshalJSON(b []byte) error {
	var m channelPollBeginEventMeta
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
	c.Choices = m.Choices
	c.BitsVoting = m.BitsVoting
	c.ChannelPointsVoting = m.ChannelPointsVoting
	c.StartedAt = m.StartedAt
	c.EndsAt = m.EndsAt

	return nil
}
