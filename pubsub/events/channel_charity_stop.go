package events

import (
	"encoding/json"
	"strconv"
	"time"
)

const CHANNEL_CHARITY_STOP_EVENT = "channel.charity_campaign.stop"

// A Channel CharityStop Event.
// Sent when a charity campaign stops
//
// Subscription name:
//
//	channel.charity_campaign.stop
//
// Required scope:
//
//	channel:read:charity
type ChannelCharityStopEvent struct {
	ID string

	BroadcasterID          int
	BroadcasterDisplayName string
	BroadcasterName        string

	CharityName        string
	CharityDescription string
	CharityLogo        string
	CharityWebsite     string
	CurrentAmount      DonationAmount
	TargetAmount       DonationAmount
	StoppedAt          time.Time
}

type channelCharityStopEventMeta struct {
	Id string `json:"id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`
	BroadcasterName        string `json:"broadcaster_user_login"`

	CharityName        string         `json:"charity_name"`
	CharityDescription string         `json:"charity_description"`
	CharityLogo        string         `json:"charity_logo"`
	CharityWebsite     string         `json:"charity_website"`
	CurrentAmount      DonationAmount `json:"current_amount"`
	TargetAmount       DonationAmount `json:"target_amount"`
	StoppedAt          time.Time      `json:"stopped_at"`
}

func (c *ChannelCharityStopEvent) UnmarshalJSON(b []byte) error {
	var m channelCharityStopEventMeta
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

	c.CharityName = m.CharityName
	c.CharityDescription = m.CharityDescription
	c.CharityLogo = m.CharityLogo
	c.CharityWebsite = m.CharityWebsite
	c.TargetAmount = m.TargetAmount
	c.CurrentAmount = m.CurrentAmount
	c.StoppedAt = m.StoppedAt

	return nil
}
