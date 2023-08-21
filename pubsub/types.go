package pubsub

// TODO: Redo probably everything

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	// ErrShardTooManyTopics returned when a shard has attempted to join too many topics
	ErrShardTooManyTopics = errors.New("too many topics on shard")
	// ErrShardIDOutOfBounds returned when an invalid shard id is provided
	ErrShardIDOutOfBounds = errors.New("shard id out of bounds")
	// ErrNonceTimeout returned when the server doesnt respond to a nonced message in time
	ErrNonceTimeout = errors.New("nonced message timeout")
	// ErrPingTimeout returned when the server takes too long to respond to a ping message
	ErrPingTimeout = errors.New("server took too long to respond to ping")

	// ErrBadMessage returned when the server receives an invalid message
	ErrBadMessage = errors.New("server received an invalid message")
	// ErrBadAuth returned when a topic doesnt have the permissions required
	ErrBadAuth = errors.New("bad authentication for topic")
	// ErrBadTopic returned when an invalid topic was requested
	ErrBadTopic = errors.New("invalid topic")
	// ErrServer returned when something went wrong on the servers end
	ErrServer = errors.New("something went wrong on the servers end")
	// ErrUnknown returned when the server sends back an error that wasnt handled by the reader
	ErrUnknown = errors.New("server sent back an unknown error")

	// ErrInvalidNonceGenerator returned when a provided nonce generator can not be used
	ErrInvalidNonceGenerator = errors.New("nonce generator is invalid")
)

// An incoming event message containing some metadata and a payload
type IncomingMessage struct {
	Metadata IncomingMessageMetadata `json:"metadata"`
	Payload  *MessagePayload         `json:"payload,omitempty"`
}

type IncomingMessageMetadata struct {
	ID                  string
	MessageType         string
	Timestamp           time.Time
	SubscriptionType    *string
	SubscriptionVersion *int
}

// Twitch sends metadata as string but but the public metadata struct uses native types. Add a private wrapper to work around it
type incomingMessageMetadataMeta struct {
	ID                  string  `json:"message_id"`
	MessageType         string  `json:"message_type"`
	Timestamp           string  `json:"message_timestamp"`
	SubscriptionType    *string `json:"subscription_type,omitempty"`
	SubscriptionVersion *string `json:"subscription_version,omitempty"`
}

type MessagePayload struct {
	Subscription SubscriptionData       `json:"subscription"`
	Event        map[string]interface{} `json:"event"`
}

type SubscriptionData struct {
	ID        string
	Status    string
	Type      string
	Version   int
	Condition SubscriptionConditionData
	Transport SubscriptionTransportData
	CreatedAt time.Time
	Cost      int
}

// And again
type subscriptionDataMeta struct {
	ID        string                    `json:"id"`
	Status    string                    `json:"status"`
	Type      string                    `json:"type"`
	Version   string                    `json:"version"`
	Condition SubscriptionConditionData `json:"condition"`
	Transport SubscriptionTransportData `json:"transport"`
	CreatedAt string                    `json:"created_at"`
	Cost      int                       `json:"cost"`
}

type SubscriptionConditionData struct {
	BroadcastUserID string `json:"broadcaster_user_id"`
}

type SubscriptionTransportData struct {
	Method    string `json:"method"`
	SessionID string `json:"session_id"`
}

// TopicData stores data about a topic
type TopicData struct {
	Topics []string `json:"topics"`
	Token  string   `json:"auth_token,omitempty"`
}

// MessageType stores the type provided in MessageData
type MessageType string

const (
	// Listen outgoing message type
	Listen MessageType = "LISTEN"
	// Unlisten outgoing message type
	Unlisten MessageType = "UNLISTEN"
	// Ping outgoing message type
	Ping MessageType = "PING"

	// Response incoming message type
	Response MessageType = "RESPONSE"
	// Message incoming message type
	Message MessageType = "MESSAGE"
	// Pong incoming message type
	Pong MessageType = "PONG"
	// Reconnect incoming message type
	Reconnect MessageType = "RECONNECT"
)

// MessageError stores the error provided in MessageData
type MessageError string

const (
	// BadMessage server received an invalid message
	BadMessage MessageError = "ERR_BADMESSAGE"
	// BadAuth provided token does not have required permissions
	BadAuth MessageError = "ERR_BADAUTH"
	// TooManyTopics attempted to listen to too many topics
	TooManyTopics MessageError = "ERR_TOO_MANY_TOPICS"
	// BadTopic provided topic is invalid
	BadTopic MessageError = "ERR_BADTOPIC"
	// InvalidTopic provided topic is invalid
	InvalidTopic MessageError = "Invalid Topic"
	// ServerError something went wrong on the servers side
	ServerError MessageError = "ERR_SERVER"
)

// NonceGenerator any function that returns a string that is different every time
type NonceGenerator func() string

// ParseTopic returns a topic string with the provided arguments
// Example: "abc", 21, [true, true, false] -> "abc.21.[true true false]"
func ParseTopic(str string, args ...interface{}) string {
	if len(args) > 0 {
		var params []string
		for _, arg := range args {
			params = append(params, fmt.Sprint(arg))
		}
		return fmt.Sprintf("%s.%s", str, strings.Join(params, "."))
	}
	return str
}

func (m *IncomingMessageMetadata) UnmarshalJSON(b []byte) error {
	var mm incomingMessageMetadataMeta
	err := json.Unmarshal(b, &mm)
	if err != nil {
		return err
	}
	m.ID = mm.ID
	m.MessageType = m.MessageType
	m.Timestamp, err = time.Parse(time.RFC3339Nano, mm.Timestamp)
	if err != nil {
		return err
	}
	m.SubscriptionType = mm.SubscriptionType
	*m.SubscriptionVersion, err = strconv.Atoi(*mm.SubscriptionVersion)
	if err != nil {
		return err
	}

	return nil
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
