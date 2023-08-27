package events

import (
	"encoding/json"
	"strconv"
)

const CHANNEL_CHEER_EVENT = "channel.cheer"

// A Channel Subscription Event.
// Sent when a user cheers on the target channel
//
// Subscription name:
//
//	channel.cheer
//
// Required scope:
//
//	bits:read
type ChannelCheerEvent struct {
	Anonymous bool

	UserID          *int
	UserName        *string
	UserDisplayName *string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Message string
	Bits    int
}

// Meta struct to convert from json full of strings to native types
type channelCheerEventMeta struct {
	Anonymous bool `json:"is_anonymous"`

	UserID          *string `json:"user_id:omitempty"`
	UserName        *string `json:"user_login:omitempty"`
	UserDisplayName *string `json:"user_name:omitempty"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	Message string `json:"message"`
	Bits    int    `json:"bits"`
}

func (c *ChannelCheerEvent) UnmarshalJSON(b []byte) error {
	var m channelCheerEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.Anonymous = m.Anonymous

	// If not anonymous, convert UserID
	if !m.Anonymous {
		tmp := 100
		tmp, err = strconv.Atoi(*m.UserID)
		c.UserID = &tmp
		if err != nil {
			return err
		}
	}
	c.UserName = m.UserName
	c.UserDisplayName = m.UserDisplayName

	c.BroadcasterID, err = strconv.Atoi(m.BroadcasterID)
	if err != nil {
		return err
	}
	c.BroadcasterName = m.BroadcasterName
	c.BroadcasterDisplayName = m.BroadcasterDisplayName

	c.Message = m.Message
	c.Bits = m.Bits

	return nil
}
