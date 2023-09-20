package events

import (
	"encoding/json"
	"strconv"
)

const CHANNEL_CHARITY_PROGRESS_EVENT = "channel.charity_campaign.progress"

// A Channel CharityProgress Event.
// Sent when a charity campaign progresses
//
// Subscription name:
//
//	channel.charity_campaign.progress
//
// Required scope:
//
//	channel:read:charity
type ChannelCharityProgressEvent struct {
	ID         string
	CampaignID string

	BroadcasterID          int
	BroadcasterDisplayName string
	BroadcasterName        string

	UserID          int
	UserName        string
	UserDisplayName string

	CharityName        string
	CharityDescription string
	CharityLogo        string
	CharityWebsite     string
	CurrentAmount      DonationAmount
	TargetAmount       DonationAmount
}

type channelCharityProgressEventMeta struct {
	Id         string `json:"id"`
	CampaignId string `json:"campaign_id"`

	BroadcasterID          string `json:"broadcaster_user_id"`
	BroadcasterDisplayName string `json:"broadcaster_user_name"`
	BroadcasterName        string `json:"broadcaster_user_login"`

	UserID          string `json:"user_id"`
	UserName        string `json:"user_login"`
	UserDisplayName string `json:"user_name"`

	CharityName        string         `json:"charity_name"`
	CharityDescription string         `json:"charity_description"`
	CharityLogo        string         `json:"charity_logo"`
	CharityWebsite     string         `json:"charity_website"`
	CurrentAmount      DonationAmount `json:"current_amount"`
	TargetAmount       DonationAmount `json:"target_amount"`
}

func (c *ChannelCharityProgressEvent) UnmarshalJSON(b []byte) error {
	var m channelCharityProgressEventMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	c.ID = m.Id
	c.CampaignID = m.CampaignId

	c.BroadcasterID, err = strconv.Atoi(m.BroadcasterID)
	if err != nil {
		return err
	}
	c.BroadcasterName = m.BroadcasterName
	c.BroadcasterDisplayName = m.BroadcasterDisplayName

	c.UserID, err = strconv.Atoi(m.UserID)
	if err != nil {
		return err
	}
	c.UserName = m.UserName
	c.UserDisplayName = m.UserDisplayName

	c.CharityName = m.CharityName
	c.CharityDescription = m.CharityDescription
	c.CharityLogo = m.CharityLogo
	c.CharityWebsite = m.CharityWebsite
	c.TargetAmount = m.TargetAmount
	c.CurrentAmount = m.CurrentAmount

	return nil
}