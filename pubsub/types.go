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

const (
	InternalServerError   = 4000 // Indicates a problem with the server (similar to an HTTP 500 status code)
	ClientSentMessage     = 4001 // Sending outgoing messages to the server is prohibited with the exception of pong messages
	FailedPing            = 4002 // You must respond to ping messages with a pong message. See https://dev.twitch.tv/docs/eventsub/websocket-reference/#ping-message.
	UnusedConnection      = 4003 // When you connect to the server, you must create a subscription within 10 seconds or the connection is closed. The time limit is subject to change
	ReconnectGraceExpired = 4004 // When you receive a session_reconnect message, you have 30 seconds to reconnect to the server and close the old connection. See https://dev.twitch.tv/docs/eventsub/websocket-reference/#reconnect-message
	NetworkTimeout        = 4005 // Transient network timeout
	NetworkError          = 4006 // Transient network error
	InvalidReconnect      = 4007 // The reconnect URL is invalid
)

// MessageType stores the type provided in MessageData
type MessageType string

// MessageError stores the error provided in MessageData
type MessageError string

// NonceGenerator any function that returns a string that is different every time
type NonceGenerator func() string

// An incoming event message containing some metadata and a payload
type IncomingMessage struct {
	Metadata IncomingMessageMetadata `json:"metadata"`
	Payload  *interface{}            `json:"payload,omitempty"`
}

type IncomingMessageMetadata struct {
	ID                  string
	MessageType         string
	Timestamp           time.Time
	SubscriptionType    *string
	SubscriptionVersion *int
}

// Twitch sends metadata as string but the public metadata struct uses native types. Add a private wrapper to work around it
type incomingMessageMetadataMeta struct {
	ID                  string  `json:"message_id"`
	MessageType         string  `json:"message_type"`
	Timestamp           string  `json:"message_timestamp"`
	SubscriptionType    *string `json:"subscription_type,omitempty"`
	SubscriptionVersion *string `json:"subscription_version,omitempty"`
}

type Subscription struct {
	ID        string
	Status    string
	Type      string
	Version   string
	Cost      string
	Condition map[string]string
	Transport map[string]string
}

type IncomingMessagePayloadSingleEvent struct {
	Subscription Subscription
	Event        interface{}
}

type IncomingMessagePayloadMultiEvent struct {
	Subscription Subscription
	Events       []interface{}
}

// TopicData stores data about a topic
type TopicData struct {
	Topics []string `json:"topics"`
	Token  string   `json:"auth_token,omitempty"`
}

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
