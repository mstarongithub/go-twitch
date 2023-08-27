package events

import (
	"encoding/json"
	"strconv"
)

// A Channel GuestStarSlotUpdate Event.
// Sent when a guest star slot is updated
//
// Subscription name:
//
//	channel.guest_star_slot.update
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
type ChannelGuestStarSlotUpdateEvent struct {
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

	SlotID           int
	HostVideoEnabled bool
	HostAudioEnabled bool
	HostVolume       int
}

// Meta struct to convert from json full of strings to native types
type channelGuestStarSlotUpdateEventMeta struct {
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

	SlotID           string `json:"slot_id"`
	HostVideoEnabled bool   `json:"host_video_enabled"`
	HostAudioEnabled bool   `json:"host_audio_enabled"`
	HostVolume       int    `json:"host_volume"`
}

func (c *ChannelGuestStarSlotUpdateEvent) UnmarshalJSON(b []byte) error {
	var m channelGuestStarSlotUpdateEventMeta
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

	c.SlotID, err = strconv.Atoi(m.SlotID)
	if err != nil {
		return err
	}
	c.HostVideoEnabled = m.HostVideoEnabled
	c.HostAudioEnabled = m.HostAudioEnabled
	c.HostVolume = m.HostVolume

	return nil
}
