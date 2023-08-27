package events

import (
	"encoding/json"
	"strconv"
)

// A Channel GuestStarGuestUpdate Event.
// Sent when a guest star guest is updated
//
// Subscription name:
//
//	channel.guest_star_guest.update
//
// Required scope:
//
//	channel:read:guest_star
//
// or
//
//	channel:manage:guest_star
//
// or
//
//	moderator:read:guest_star
//
// or
//
//	moderator:manage:guest_star
//
// # NOTE: Public Beta Event
type ChannelGuestStarGuestUpdateEvent struct {
	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	SessionID string

	ModID          int
	ModName        string
	ModDisplayName string

	GuestID          int
	GuestName        string
	GuestDisplayName string
	Slot             int
	State            string
}

// Meta struct to convert from json full of strings to native types
type channelGuestStarGuestUpdateEventMeta struct {
	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	SessionID string `json:"session_id"`

	ModID          string `json:"moderator_user_id"`
	ModName        string `json:"moderator_user_name"`
	ModDisplayName string `json:"moderator_user_login"`

	GuestID          string `json:"guest_user_id"`
	GuestName        string `json:"guest_user_name"`
	GuestDisplayName string `json:"guest_user_login"`

	Slot  string `json:"slot_id"`
	State string `json:"state"`
}

func (c *ChannelGuestStarGuestUpdateEvent) UnmarshalJSON(b []byte) error {
	var m channelGuestStarGuestUpdateEventMeta
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

	c.SessionID = m.SessionID

	c.ModID, err = strconv.Atoi(m.ModID)
	if err != nil {
		return err
	}
	c.ModName = m.ModName
	c.ModDisplayName = m.ModDisplayName

	c.GuestID, err = strconv.Atoi(m.GuestID)
	if err != nil {
		return err
	}
	c.GuestName = m.GuestName
	c.GuestDisplayName = m.GuestDisplayName

	c.Slot, err = strconv.Atoi(m.Slot)
	if err != nil {
		return err
	}
	c.State = m.State

	return nil
}
