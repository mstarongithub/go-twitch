package events

import (
	"encoding/json"
	"strconv"
	"time"
)

type RevocationPayload struct {
	Subscription RevocationSubscription `json:"subscription"`
}

type RevocationSubscription struct {
	ID        string
	Status    string
	Type      string
	Version   string
	Cost      int
	Condition interface{}
	Transport TransportData
	CreatedAt time.Time
}

type revocationSubscriptionMeta struct {
	ID        string
	Status    string
	Type      string
	Version   string
	Cost      string
	Condition interface{}
	Transport TransportData
	CreatedAt string
}

func (s *RevocationSubscription) UnmarshalJSON(b []byte) error {
	var m revocationSubscriptionMeta
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	s.ID = m.ID
	s.Status = m.Status
	s.Type = m.Type
	s.Version = m.Version
	s.Cost, err = strconv.Atoi(m.Cost)
	if err != nil {
		return err
	}
	s.Condition = m.Condition
	s.Transport = m.Transport
	s.CreatedAt, err = time.Parse(time.RFC3339Nano, m.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
