package events

import (
	"encoding/json"
	"strconv"
	"time"
)

// Payload of a notification message
type NotificationPayload struct {
	Subscription SubscriptionData `json:"subscription"`
	Event        *interface{}     `json:"event"`
}

type SubscriptionData struct {
	ID        string
	Status    string
	Type      string
	Version   int
	Condition *interface{}
	Transport TransportData
	CreatedAt time.Time
	Cost      int
}

// And again
type subscriptionDataMeta struct {
	ID        string        `json:"id"`
	Status    string        `json:"status"`
	Type      string        `json:"type"`
	Version   string        `json:"version"`
	Condition *interface{}  `json:"condition,omitempty"`
	Transport TransportData `json:"transport"`
	CreatedAt string        `json:"created_at"`
	Cost      int           `json:"cost"`
}

func (s *SubscriptionData) UnmarshalJSON(b []byte) error {
	var m subscriptionDataMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	s.ID = m.ID
	s.Status = m.Status
	s.Type = m.Type
	s.Version, err = strconv.Atoi(m.Version)
	if err != nil {
		return err
	}
	s.Condition = m.Condition
	s.Transport = m.Transport
	s.CreatedAt, err = time.Parse(time.RFC3339Nano, m.CreatedAt)
	if err != nil {
		return err
	}
	s.Cost = m.Cost

	return nil
}
