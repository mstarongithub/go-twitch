package events

import (
	"encoding/json"
	"strconv"
)

// A Channel GuestStarSettingsUpdate Event.
// Sent when a guest star setting is updated
//
// Subscription name:
//
//	channel.guest_star_settings.update
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
type ChannelGuestStarSettingsUpdateEvent struct {
	BroadcasterID          int
	BroadcasterName        string
	BroadcasterDisplayName string

	IsModeratorSendLiveEnabled  bool
	SlotCount                   int
	IsBrowserSourceAudioEnabled bool
	GroupLayout                 string
}

// Meta struct to convert from json full of strings to native types
type channelGuestStarSettingsUpdateEventMeta struct {
	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterName        string `json:"broadcaster_user_login"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`

	IsModeratorSendLiveEnabled  bool   `json:"is_moderator_send_live_enabled"`
	SlotCount                   int    `json:"slot_count"`
	IsBrowserSourceAudioEnabled bool   `json:"is_browser_source_audio_enabled"`
	GroupLayout                 string `json:"group_layout"`
}

func (c *ChannelGuestStarSettingsUpdateEvent) UnmarshalJSON(b []byte) error {
	var m channelGuestStarSettingsUpdateEventMeta
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

	c.IsModeratorSendLiveEnabled = m.IsModeratorSendLiveEnabled
	c.SlotCount = m.SlotCount
	c.IsBrowserSourceAudioEnabled = m.IsBrowserSourceAudioEnabled
	c.GroupLayout = m.GroupLayout

	return nil
}
