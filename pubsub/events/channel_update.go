package events

import (
	"encoding/json"
	"strconv"
)

// A Channel Update Event.
//
// Sent when the category, title, labels or language of a stream updates.
//
// Subscription name:
//
//	channel.update
//
// Required scope:
//
//	None
type ChannelUpdateEvent struct {
	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Title        string
	Language     string
	CategoryID   int
	CategoryName string
	Labels       []string
}

// Meta struct to convert from json full of strings to native types
type channelUpdateEventMeta struct {
	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Title        string   `json:"title"`
	Language     string   `json:"language"`
	CategoryID   string   `json:"category_id"`
	CategoryName string   `json:"category_name"`
	Labels       []string `json:"content_classification_labels"`
}

func (c *ChannelUpdateEvent) UnmarshalJSON(b []byte) error {
	var m channelUpdateEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.BroadcasterID, err = strconv.Atoi(m.BroadcasterID)
	if err != nil {
		return err
	}
	c.BroadcasterName = m.BroadcasterName
	c.BroadcasterDisplayName = m.BroadcasterDisplayName

	c.Title = m.Title
	c.Language = m.Language
	c.CategoryID, err = strconv.Atoi(m.CategoryID)
	if err != nil {
		return err
	}
	c.CategoryName = m.CategoryName
	c.Labels = m.Labels

	return nil
}
