package eventsub

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	messages "github.com/Adeithe/go-twitch/eventsub/messages"
	"github.com/Adeithe/go-twitch/eventsub/nonce"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Conn stores data about a PubSub connection
type Conn struct {
	length int
	socket *websocket.Conn
	done   chan bool

	isConnected bool
	latency     time.Duration
	ping        chan bool

	generator NonceGenerator
	topics    map[string][]Topic
	pending   map[string]chan error
	nonces    sync.Mutex
	listeners sync.Mutex
	writer    sync.Mutex

	onRawMessage   []func([]byte)
	onNotification []func(IncomingMessage)
	onPong         []func(time.Duration)
	onReconnect    []func()
	onDisconnect   []func()
	onError        []func(err error)

	// Store those for reconnects
	scheme, host, path string

	idChecker idChecker

	lastMessage      time.Time
	keepAliveTimeout time.Duration // Max timeout diff in seconds
	sessionID        string
	cleanRejoin      bool
}

// IP for the PubSub server
const IP = "wss://eventsub-edge.twitch.tv"

// Connect to the default Twitch server
func (conn *Conn) Connect() error {
	return conn.ConnectCustomServerRawURL(IP)
}

// Connect to a custom server using a raw URL
func (conn *Conn) ConnectCustomServerRawURL(raw string) error {
	u, err := decodeURL(raw)
	if err != nil {
		return err
	}
	return conn.ConnectCustomServer(u.Scheme, u.Host, u.Path)
}

// Connect to a custom server
// Useful for testing with, for example, the Twitch CLI
func (conn *Conn) ConnectCustomServer(scheme, host, path string) error {
	conn.scheme = scheme
	conn.host = host
	conn.path = path

	u := url.URL{Scheme: scheme, Host: host, Path: path}
	socket, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	if conn.length < 1 {
		conn.length = 50
	}
	if conn.generator == nil {
		conn.generator = nonce.WichmannHill
	}
	conn.socket = socket
	conn.done = make(chan bool)
	conn.isConnected = true
	go conn.reader()
	go conn.timeoutChecker()
	if conn.topics != nil && !conn.cleanRejoin {
		var wg sync.WaitGroup
		conn.listeners.Lock()
		rejoined := make(map[string][]Topic)
		for token, topics := range conn.topics {
			wg.Add(1)
			go func(token string, topics ...Topic) {
				for _, topic := range topics {
					if err := conn.ListenWithAuth(topic.ChannelID, token, topic.Name, topic.Version); err == nil {
						rejoined[token] = topics
					}
				}
				wg.Done()
			}(token, topics...)
		}
		conn.listeners.Unlock()
		wg.Wait()
		conn.listeners.Lock()
		defer conn.listeners.Unlock()
		conn.topics = rejoined
	}
	return nil
}

// Reconnect to the PubSub server
func (conn *Conn) Reconnect() error {
	if conn.isConnected {
		conn.Close()
	}
	if err := conn.ConnectCustomServer(conn.scheme, conn.host, conn.path); err != nil {
		return err
	}
	for _, f := range conn.onReconnect {
		go f()
	}
	return nil
}

// Keep for clean closing of connection
func (conn *Conn) write(msgType int, data []byte) error {
	conn.writer.Lock()
	defer conn.writer.Unlock()
	return conn.socket.WriteMessage(msgType, data)
}

// Close the connection to the PubSub server
func (conn *Conn) Close() {
	conn.write(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	timer := time.NewTimer(time.Second)
	defer timer.Stop()
	select {
	case <-conn.done:
	case <-timer.C:
		conn.socket.Close()
	}
}

// IsConnected returns true if the socket is actively connected
func (conn *Conn) IsConnected() bool {
	return conn.isConnected
}

// SetMaxTopics changes the maximum number of topics the connection can listen to
func (conn *Conn) SetMaxTopics(max int) {
	if max < 1 {
		max = 50
	}
	conn.length = max
}

// GetNumTopics returns the number of topics the connection is actively listening to
func (conn *Conn) GetNumTopics() (n int) {
	conn.listeners.Lock()
	defer conn.listeners.Unlock()
	if conn.topics != nil {
		for _, topics := range conn.topics {
			n += len(topics)
		}
	}
	return
}

// HasTopic returns true if the connection is actively listening to the provided topic
func (conn *Conn) HasTopic(topic Topic) bool {
	conn.listeners.Lock()
	defer conn.listeners.Unlock()
	for _, g := range conn.topics {
		for _, t := range g {
			if topic.Name == t.Name && topic.Version == t.Version && topic.ChannelID == topic.ChannelID {
				return true
			}
		}
	}
	return false
}

// Listen to a topic using no authentication token
//
// This operation will block, giving the server up to 5 seconds to respond after correcting for latency before failing
func (conn *Conn) Listen(channelID int, topic string, version int) error {
	return conn.ListenWithAuth(channelID, "", topic, version)
}

// ListenWithAuth starts listening to a topic using the provided authentication token
//
// This operation will block, giving the server up to 5 seconds to respond after correcting for latency before failing
func (conn *Conn) ListenWithAuth(channelID int, token string, topic string, version int) error {
	if conn.GetNumTopics()+1 > conn.length {
		return ErrShardTooManyTopics
	}

	jsBody := []byte(fmt.Sprintf(""))
	http.NewRequest("POST", "https://api.twitch.tv/helix/eventsub/subscriptions", bytes.NewReader(jsBody))
	conn.listeners.Lock()
	defer conn.listeners.Unlock()
	if conn.topics == nil {
		conn.topics = make(map[string][]string)
	}
	conn.topics[token] = append(conn.topics[token], topic)
	return nil
}

// Unlisten from the provided topics
//
// This operation will block, giving the server up to 5 seconds to respond after correcting for latency before failing
func (conn *Conn) Unlisten(topics ...Topic) error {
	var unlisten []Topic
	for _, topic := range topics {
		if conn.HasTopic(topic) {
			unlisten = append(unlisten, topic)
		}
	}
	if len(unlisten) < 1 {
		return nil
	}
	conn.listeners.Lock()
	for token, topics := range conn.topics {
		var newTopics []Topic
		for _, topic := range topics {
			var b bool
			for _, t := range unlisten {
				if topic.Name == t.Name && topic.Version == t.Version && topic.ChannelID == topic.ChannelID {
					b = true
					break
				}
			}
			if !b {
				newTopics = append(newTopics, topic)
			}
		}
		conn.topics[token] = newTopics
	}
	conn.listeners.Unlock()
	// TODO: Unlisten to topics here
	return nil
}

// OnMessage event called after a message is receieved
func (conn *Conn) OnRawMessage(f func([]byte)) {
	conn.onRawMessage = append(conn.onRawMessage, f)
}

func (conn *Conn) OnNotification(f func(IncomingMessage)) {
	conn.onNotification = append(conn.onNotification, f)
}

// OnPong event called after a Pong message is received, updating the latency
func (conn *Conn) OnPong(f func(time.Duration)) {
	conn.onPong = append(conn.onPong, f)
}

// OnReconnect event called after the connection is reopened
func (conn *Conn) OnReconnect(f func()) {
	conn.onReconnect = append(conn.onReconnect, f)
}

// OnDisconnect event called after the connection is closed
func (conn *Conn) OnDisconnect(f func()) {
	conn.onDisconnect = append(conn.onDisconnect, f)
}

// OnError event called if an error is encountered that can't be easily recovered from
func (conn *Conn) OnError(f func(error)) {
	conn.onError = append(conn.onError, f)
}

func (conn *Conn) reader() {
	for {
		msgType, data, err := conn.socket.ReadMessage()
		if err != nil || msgType == websocket.CloseMessage {
			break
		}

		// First send the message to any raw receivers
		for _, f := range conn.onRawMessage {
			// Don't send the original package, make a copy first
			// Rather use more memory than have race conditions on read
			bytesCopy := make([]byte, len(data))
			copy(bytesCopy, data)
			f(bytesCopy)
		}

		var msg IncomingMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			continue
		}

		if conn.idChecker.Has(msg.Metadata.ID) {
			continue
		}
		conn.idChecker.Add(msg.Metadata.ID)
		conn.lastMessage = msg.Metadata.Timestamp

		switch msg.Metadata.MessageType {
		case WelcomeMessage:
			conn.handleWelcomeMessage(msg)
		case KeepaliveMessage:
			continue
		case NotificationMessage:
			conn.handleMessage(msg)
		case ReconnectMessage:
			conn.handleReconnectRequest(msg)
		case RevocationMessage:
			conn.handleRevocation(msg)
		}

	}
	conn.socket.Close()
	conn.isConnected = false
	close(conn.done)
	for _, f := range conn.onDisconnect {
		go f()
	}
}

// just keep checking if the server fails to respond within the timeout
func (conn *Conn) timeoutChecker() {
	initialTimeout := time.Minute * 5
	timer := time.NewTimer(initialTimeout)
	defer timer.Stop()
	for {
		select {
		case <-conn.done:
			return
		case <-timer.C:
			if conn.keepAliveTimeout == 0 {
				// Didn't get a welcome message until timeout. Close
				conn.Close()
			} else {
				if time.Now().Sub(conn.lastMessage).Abs() > conn.keepAliveTimeout {
					// No message received within timeout. Close
					conn.Close()
					for _, f := range conn.onError {
						f(errors.New("conn.timeoutChecker: Twitch endpoint failed to reply in time. Closing connection"))
					}
				} else {
					// Get diff between now and last message, then subtract it from the timeout to get the new time to wait
					// Done to wait from now until lastMessage + timeout
					diff := time.Now().Sub(conn.lastMessage).Abs()
					timer.Reset(conn.keepAliveTimeout - diff)
				}
			}
		}
	}
}

func (conn *Conn) handleMessage(i IncomingMessage) {
	var typed messages.NotificationPayload
	err := mapstructure.Decode(i.Payload, &typed)
	if err != nil {
		for _, f := range conn.onError {
			f(fmt.Errorf("conn.handleMessage: Failed to cast Payload to correct type: %w", err))
		}
		return
	}

	for _, f := range conn.onNotification {
		f(i)
	}
}

func (conn *Conn) handleReconnectRequest(msg IncomingMessage) {
	var typed messages.ReconnectPayload
	err := mapstructure.Decode(msg.Payload, &typed)

	conn.Close()
	// Handle err from Decode after close
	if err != nil {
		for _, f := range conn.onError {
			f(fmt.Errorf("conn.handleReconnectRequest: Failed to cast Payload to correct type: %w", err))
		}
		return
	}

	// Store old url first
	old := url.URL{Scheme: conn.scheme, Host: conn.host, Path: conn.path}
	// set cleanRejoin to true to indicate that we should attempt a clean rejoin first
	conn.cleanRejoin = true
	err = conn.ConnectCustomServerRawURL(typed.Session.ReconnectURL)
	if err != nil {
		// Failed a clean rejoin. Flag it
		conn.cleanRejoin = false
		// Try to reconnect to old target and request topics again
		err := conn.ConnectCustomServer(old.Scheme, old.Host, old.Path)
		if err != nil {
			// Just give up at this point, send the error out and return
			for _, f := range conn.onError {
				f(fmt.Errorf("conn.handleReconnectRequest: Failed to connect to old endpoint after new endpoint failed: %w", err))
			}
		}
	}
}

// No idea how to deal with revocations
// Shouldn't have to either in hopes of people not running something with this module long enough for it to matter
func (conn *Conn) handleRevocation(msg IncomingMessage) {
	var typed messages.RevocationPayload
	err := mapstructure.Decode(msg.Payload, &typed)
	if err != nil {
		for _, f := range conn.onError {
			f(fmt.Errorf("conn.handleRevocation: Failed to cast Payload to correct type: %w", err))
		}
	}
	switch typed.Subscription.Status {
	case "authorization_revoked":
	case "user_removed":
	case "version_removed":
	default:
		for _, f := range conn.onError {
			f(errors.New("unknown revocation message"))
		}
	}
}

func (conn *Conn) handleWelcomeMessage(msg IncomingMessage) {
	var typed messages.WelcomePayload
	err := mapstructure.Decode(msg.Payload, &typed)
	if err != nil {
		for _, f := range conn.onError {
			f(fmt.Errorf("conn.handleWelcomeMessage: Failed to cast Payload to correct type: %w", err))
			return
		}
	}
	conn.keepAliveTimeout = time.Duration(typed.Session.KeepaliveTimeout)
	conn.sessionID = typed.Session.ID
}
