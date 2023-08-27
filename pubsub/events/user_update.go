package events

import (
	"encoding/json"
	"strconv"
)

// A User Update Event.
//
// Sent when a user updates their account.
//
// Subscription name:
//
//	user.update
//
// Required scope:
//
//	None
//
// Optional scope (for user email):
//
//	user:read:email
type UserUpdateEvent struct {
	UserID          int
	UserName        string
	UserDisplayName string

	Email         *string
	EmailVerified bool
	Description   string
}

// Meta struct to convert from json full of strings to native types
type userUpdateEventMeta struct {
	UserID          string `json:"user_id"`
	UserName        string `json:"user_login"`
	UserDisplayName string `json:"user_name"`

	Email         *string `json:"email,omitempty"`
	EmailVerified bool    `json:"email_verified"`
	Description   string  `json:"description"`
}

func (c *UserUpdateEvent) UnmarshalJSON(b []byte) error {
	var m userUpdateEventMeta
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

	c.Email = m.Email
	c.EmailVerified = m.EmailVerified
	c.Description = m.Description

	return nil
}
