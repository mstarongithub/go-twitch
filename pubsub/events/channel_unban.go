package events

import (
	"encoding/json"
	"strconv"
)

// A Channel Unban Event.
// Sent when a user gets unbanned
//
// Subscription name:
//
//	channel.unban
//
// Required scope:
//
//	channel:moderate
type ChannelUnbanEvent struct {
	UserID          int
	UserName        string
	UserDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	ModID          int
	ModName        string
	ModDisplayName string
}

// Meta struct to convert from json full of strings to native types
type channelUnbanEventMeta struct {
	UserID          string `json:"user_id"`
	UserName        string `json:"user_login"`
	UserDisplayName string `json:"user_name"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	ModID          string `json:"moderator_user_id"`
	ModName        string `json:"moderator_user_login"`
	ModDisplayName string `json:"moderator_user_name"`
}

func (c *ChannelUnbanEvent) UnmarshalJSON(b []byte) error {
	var m channelUnbanEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.UserID, err = strconv.Atoi(m.UserID)
	if err != nil {
		return err
	}
	c.UserName = m.UserName
	c.UserDisplayName = m.UserDisplayName

	c.BroadcasterID, err = strconv.Atoi(m.BroadcasterID)
	if err != nil {
		return err
	}
	c.BroadcasterName = m.BroadcasterName
	c.BroadcasterDisplayName = m.BroadcasterDisplayName

	c.ModID, err = strconv.Atoi(m.ModID)
	if err != nil {
		return err
	}
	c.ModName = m.ModName
	c.ModDisplayName = m.ModDisplayName

	return nil
}
