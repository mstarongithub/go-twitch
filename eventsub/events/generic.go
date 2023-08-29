package events

import (
	"encoding/json"
	"strconv"
)

// An emote inside of a message
type MessageEmote struct {
	Begin int    `json:"begin"`
	End   int    `json:"end"`
	ID    string `json:"id"`
}

// An event message
type EventMessage struct {
	Text   string         `json:"text"`
	Emotes []MessageEmote `json:"emotes"`
}

// A poll choice option.
//
// BitVotes, ChannelPointVotes and Votes are not set on add
type PollChoice struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	BitVotes          *int   `json:"bit_votes,omitempty"`
	ChannelPointVotes *int   `json:"channel_points_votes,omitempty"`
	Votes             *int   `json:"votes,omitempty"`
}

type PredictionOutcome struct {
	Id            string      `json:"id"`
	Title         string      `json:"title"`
	Color         string      `json:"color"`
	ChannelPoints *int        `json:"channel_points,omitempty"`
	TopPredictors []Predictor `json:"top_predictors,omitempty"` // Max 10 entries
}

// A user participating in a prediction
type Predictor struct {
	DisplayName       string
	Name              string
	ID                int
	ChannelPointsWon  *int
	ChannelPointsUsed int
}

type predictorMeta struct {
	DisplayName       string `json:"user_name"`
	Name              string `json:"user_login"`
	ID                string `json:"user_id"`
	ChannelPointsWon  *int   `json:"channel_points_won,omitempty"`
	ChannelPointsUsed int    `json:"channel_points_used"`
}

type HypeTrainContribution struct {
	UserId    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
	Type      string `json:"type"`
	Total     int    `json:"total"`
}

type DonationAmount struct {
	Value         int    `json:"value"`
	DecimalPlaces int    `json:"decimal_places"`
	Currency      string `json:"currency"`
}

func (c *Predictor) UnmarshalJSON(b []byte) error {
	var m predictorMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.ID, err = strconv.Atoi(m.ID)
	if err != nil {
		return err
	}
	c.Name = m.Name
	c.DisplayName = m.DisplayName

	return nil
}
