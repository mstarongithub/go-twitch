package events

import (
	"encoding/json"
	"strconv"
)

const CHANNEL_RAID_EVENT = "channel.raid"

// A Channel Raid Event.
// Sent when a target channel gets raided
//
// Subscription name:
//
//	channel.raid
//
// Required scope:
//
//	None
type ChannelRaidEvent struct {
	FromUserID          int
	FromUserName        string
	FromUserDisplayName string

	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	Viewers int
}

// Meta struct to convert from json full of strings to native types
type channelRaidEventMeta struct {
	FromUserID          string `json:"from_broadcaster_user_id"`
	FromUserName        string `json:"from_broadcaster_user_login"`
	FromUserDisplayName string `json:"from_broadcaster_user_name"`

	BroadcasterID          string `json:"to_broadcaster_user_id"`
	BroadcasterName        string `json:"to_broadcaster_user_login"`
	BroadcasterDisplayName string `json:"to_broadcaster_user_name"`

	Viewers int `json:"viewers"`
}

func (c *ChannelRaidEvent) UnmarshalJSON(b []byte) error {
	var m channelRaidEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.FromUserID, err = strconv.Atoi(m.FromUserID)
	if err != nil {
		return err
	}
	c.FromUserName = m.FromUserName
	c.FromUserDisplayName = m.FromUserDisplayName

	c.BroadcasterID, err = strconv.Atoi(m.BroadcasterID)
	if err != nil {
		return err
	}
	c.BroadcasterName = m.BroadcasterName
	c.BroadcasterDisplayName = m.BroadcasterDisplayName

	c.Viewers = m.Viewers

	return nil
}
