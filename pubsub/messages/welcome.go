package events

import (
	"encoding/json"
	"time"
)

// The payload of a welcome message
type WelcomePayload struct {
	Session WelcomeSession `json:"session"` // It's just a wrapper around a WelcomeSession
}

type WelcomeSession struct {
	ID               string
	Status           string
	KeepaliveTimeout int // Max amount of seconds before sending an event subscription after welcome / for the empty time before a keepalive message
	ReconnectURL     string
	ConnectedAt      time.Time
}

type welcomeSessionMeta struct {
	ID               string `json:"id"`
	Status           string `json:"status"`
	KeepaliveTimeout int    `json:"keepalive_timeout_seconds"`
	ReconnectURL     string `json:"reconnect_url"`
	ConnectedAt      string `json:"connected_at"`
}

func (s *WelcomeSession) UnmarshalJSON(b []byte) error {
	var m welcomeSessionMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	s.ID = m.ID
	s.Status = m.Status
	s.KeepaliveTimeout = m.KeepaliveTimeout
	s.ReconnectURL = m.ReconnectURL
	s.ConnectedAt, err = time.Parse(time.RFC3339Nano, m.ConnectedAt)
	if err != nil {
		return err
	}

	return nil
}
