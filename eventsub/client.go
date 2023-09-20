package eventsub

import (
	"fmt"
	"sync"
	"time"

	"github.com/Adeithe/go-twitch/eventsub/events"
)

// Client stores data about a PubSub shard manager
type Client struct {
	length        int
	topicsLength  int
	shards        map[int]*Conn
	awaitingClose int

	onShardConnect       []func(int)
	onShardRawMessage    []func(int, []byte)
	onShardLatencyUpdate []func(int, time.Duration)
	onShardReconnect     []func(int)
	onShardDisconnect    []func(int)
	onError              []func(error)

	mx sync.Mutex
	wg sync.WaitGroup

	handlerStore
}

var _ = &Client{}

// New PubSub Client
//
// The client uses a sharding system to comply with limits as listed on the Twitch PubSub Documentation.
// Twitch recommends no more than 10 simultaneous shards and no more than 50 topics per shard. These are set by default.
//
// If for any reason a shard attempts to listen to more topics than the server allows, it will attempt to correct for it.
//
// See: https://dev.twitch.tv/docs/pubsub
func New() *Client {
	return &Client{length: 10, topicsLength: 50}
}

// SetMaxShards set the maximum number of shards
//
// Default: 10
func (client *Client) SetMaxShards(max int) {
	if max < 1 {
		max = 10
	}
	client.length = max
}

// SetMaxTopicsPerShard set the maximum number of topics for each shard
//
// Default: 50
func (client *Client) SetMaxTopicsPerShard(max int) {
	if max < 1 {
		max = 50
	}
	client.mx.Lock()
	defer client.mx.Unlock()
	client.topicsLength = max
	for _, shard := range client.shards {
		shard.SetMaxTopics(client.topicsLength)
	}
}

// GetNumShards returns the number of active shards
func (client *Client) GetNumShards() int {
	return len(client.shards)
}

// GetNumTopics returns the number of topics being listened to across all shards
func (client *Client) GetNumTopics() (n int) {
	client.mx.Lock()
	defer client.mx.Unlock()
	for _, shard := range client.shards {
		n += shard.GetNumTopics()
	}
	return
}

// GetNextShard returns the first shard that can accept topics
func (client *Client) GetNextShard() (*Conn, error) {
	client.mx.Lock()
	shardID := len(client.shards)
	for id, conn := range client.shards {
		if conn.GetNumTopics() < conn.length {
			shardID = id
			break
		}
	}
	client.mx.Unlock()
	return client.GetShard(shardID)
}

// GetShard retrieves or creates a shard based on the provided id
func (client *Client) GetShard(id int) (*Conn, error) {
	if client.length < 1 {
		client.length = 10
	}
	if client.topicsLength < 1 {
		client.topicsLength = 50
	}
	if id < 0 || id > client.length-1 {
		return nil, ErrShardIDOutOfBounds
	}
	client.mx.Lock()
	defer client.mx.Unlock()
	if client.shards == nil {
		client.shards = make(map[int]*Conn)
	}
	if client.shards[id] == nil {
		conn := &Conn{length: client.topicsLength}
		conn.OnRawMessage(func(data []byte) {
			for _, f := range client.onShardRawMessage {
				go f(id, data)
			}
		})
		conn.OnNotification(client.handleMessage)
		conn.OnPong(func(latency time.Duration) {
			for _, f := range client.onShardLatencyUpdate {
				go f(id, latency)
			}
		})
		conn.OnReconnect(func() {
			for _, f := range client.onShardReconnect {
				go f(id)
			}
		})
		conn.OnDisconnect(func() {
			client.mx.Lock()
			defer client.mx.Unlock()
			for _, f := range client.onShardDisconnect {
				go f(id)
			}
			if client.awaitingClose > 0 {
				client.awaitingClose--
				delete(client.shards, id)
				client.wg.Done()
			}
		})
		client.shards[id] = conn
	}
	shard := client.shards[id]
	if !shard.IsConnected() {
		if err := shard.Connect(); err != nil {
			return nil, err
		}
		for _, f := range client.onShardConnect {
			go f(id)
		}
	}
	return shard, nil
}

// Close all active shards
func (client *Client) Close() {
	client.mx.Lock()
	for _, shard := range client.shards {
		client.wg.Add(1)
		client.awaitingClose++
		shard.Close()
	}
	client.mx.Unlock()
	client.wg.Wait()
	client.awaitingClose = 0
}

// Listen to a topic on the best available shard
func (client *Client) Listen(topic Topic) error {
	shard, err := client.GetNextShard()
	if err != nil {
		return err
	}
	client.mx.Lock()
	defer client.mx.Unlock()
	if err := shard.Listen(topic); err != nil {
		if err == ErrShardTooManyTopics {
			client.SetMaxTopicsPerShard(shard.GetNumTopics())
			shard, err := client.GetNextShard()
			if err != nil {
				return err
			}
			return shard.Listen(topic)
		}
		return err
	}
	return nil
}

// ListenWithAuth starts listening to a topic on the best available shard using the provided authentication token
func (client *Client) ListenWithAuth(token string, topic Topic) error {
	shard, err := client.GetNextShard()
	if err != nil {
		return err
	}
	client.mx.Lock()
	defer client.mx.Unlock()
	if err := shard.ListenWithAuth(token, topic); err != nil {
		if err == ErrShardTooManyTopics {
			client.SetMaxTopicsPerShard(shard.GetNumTopics())
			shard, err := client.GetNextShard()
			if err != nil {
				return err
			}
			return shard.ListenWithAuth(token, topic)
		}
		return err
	}
	return nil
}

// Unlisten from the provided topics
//
// Will return the first error that occurs, if any
func (client *Client) Unlisten(topics ...Topic) error {
	client.mx.Lock()
	defer client.mx.Unlock()
	for _, shard := range client.shards {
		for _, topic := range topics {
			if shard.HasTopic(topic) {
				if err := shard.Unlisten(topic); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// OnShardConnect event called after a shard connects to the PubSub server
func (client *Client) OnShardConnect(f func(int)) {
	client.onShardConnect = append(client.onShardConnect, f)
}

// OnShardMessage event called after a shard gets a PubSub message
func (client *Client) OnShardMessage(f func(int, []byte)) {
	client.onShardRawMessage = append(client.onShardRawMessage, f)
}

// OnShardLatencyUpdate event called after a shards latency is updated
func (client *Client) OnShardLatencyUpdate(f func(int, time.Duration)) {
	client.onShardLatencyUpdate = append(client.onShardLatencyUpdate, f)
}

// OnShardReconnect event called after a shard reconnects to the PubSub server
func (client *Client) OnShardReconnect(f func(int)) {
	client.onShardReconnect = append(client.onShardReconnect, f)
}

// OnShardDisconnect event called after a shard is disconnected from the PubSub server
func (client *Client) OnShardDisconnect(f func(int)) {
	client.onShardDisconnect = append(client.onShardDisconnect, f)
}

func (client *Client) handleMessage(i IncomingMessage) {
	event, err := events.ConvertMapToEvent(i.Payload.(map[string]interface{}), *i.Metadata.SubscriptionType)
	if err != nil {
		for _, f := range client.onError {
			f(fmt.Errorf("client.handleMessage: Failed to cast Payload to valid type: %w", err))
			return
		}
	}
	// I **HATE** that I have to manually switch on every type. Doesn't even matter if I go for the type string or actual type
	switch event.(type) {
	case events.ChannelBanEvent:
		client.sendChannelBan(event.(events.ChannelBanEvent))
	case events.ChannelCharityDonateEvent:
		client.sendChannelCharityDonate(event.(events.ChannelCharityDonateEvent))
	case events.ChannelCharityProgressEvent:
		client.sendChannelCharityProgress(event.(events.ChannelCharityProgressEvent))
	case events.ChannelCharityStartEvent:
		client.sendChannelCharityProgress(event.(events.ChannelCharityProgressEvent))
	case events.ChannelCharityStopEvent:
		client.sendChannelCharityStop(event.(events.ChannelCharityStopEvent))
	case events.ChannelCheerEvent:
		client.sendChannelCheer(event.(events.ChannelCheerEvent))
	case events.ChannelFollowEvent:
		client.sendChannelFollow(event.(events.ChannelFollowEvent))
	case events.ChannelGoalBeginEvent:
		client.sendChannelGoalBegin(event.(events.ChannelGoalBeginEvent))
	case events.ChannelGoalEndEvent:
		client.sendChannelGoalEnd(event.(events.ChannelGoalEndEvent))
	case events.ChannelGoalProgressEvent:
		client.sendChannelGoalProgress(event.(events.ChannelGoalProgressEvent))
	case events.ChannelGuestStarGuestUpdateEvent:
		client.sendChannelGuestStarGuestGuestUpdate(event.(events.ChannelGuestStarGuestUpdateEvent))
	case events.ChannelGuestStarSessionBeginEvent:
		client.sendChannelGuestStarSessionBegin(event.(events.ChannelGuestStarSessionBeginEvent))
	case events.ChannelGuestStarSessionEndEvent:
		client.sendChannelGuestStarSessionEnd(event.(events.ChannelGuestStarSessionEndEvent))
	case events.ChannelGuestStarSettingsUpdateEvent:
		client.sendChannelGuestStarSettingsUpdate(event.(events.ChannelGuestStarSettingsUpdateEvent))
	case events.ChannelGuestStarSlotUpdateEvent:
		client.sendChannelGuestStarSlotUpdate(event.(events.ChannelGuestStarSlotUpdateEvent))
	case events.ChannelHypeTrainBeginEvent:
		client.sendChannelHypeTrainBegin(event.(events.ChannelHypeTrainBeginEvent))
	case events.ChannelHypeTrainEndEvent:
		client.sendChannelHypeTrainEnd(event.(events.ChannelHypeTrainEndEvent))
	case events.ChannelHypeTrainProgressEvent:
		client.sendChannelHypeTrainProgress(event.(events.ChannelHypeTrainProgressEvent))
	case events.ChannelModPromotionEvent:
		client.sendChannelModeratorAdd(event.(events.ChannelModPromotionEvent))
	case events.ChannelModDemotionEvent:
		client.sendChannelModeratorRemove(event.(events.ChannelModDemotionEvent))
	case events.ChannelPointsRedemptionAddEvent:
		client.sendChannelPointsRedemptionAdd(event.(events.ChannelPointsRedemptionAddEvent))
	case events.ChannelPointsRedemptionUpdateEvent:
		client.sendChannelPointsRedemptionUpdate(event.(events.ChannelPointsRedemptionUpdateEvent))
	case events.ChannelPointRewardsAddEvent:
		client.sendChannelPointsRewardsAdd(event.(events.ChannelPointRewardsAddEvent))
	case events.ChannelPointRewardsRemoveEvent:
		client.sendChannelPointsRewardsRemove(event.(events.ChannelPointRewardsRemoveEvent))
	case events.ChannelPointRewardsUpdateEvent:
		client.sendChannelPointsRewardsUpdate(event.(events.ChannelPointRewardsUpdateEvent))
	case events.ChannelPollBeginEvent:
		client.sendChannelPollBegin(event.(events.ChannelPollBeginEvent))
	case events.ChannelPollEndEvent:
		client.sendChannelPollEnd(event.(events.ChannelPollEndEvent))
	case events.ChannelPollProgressEvent:
		client.sendChannelPollProgress(event.(events.ChannelPollProgressEvent))
	case events.ChannelPredictionBeginEvent:
		client.sendChannelPredictionBegin(event.(events.ChannelPredictionBeginEvent))
	case events.ChannelPredictionEndEvent:
		client.sendChannelPredictionEnd(event.(events.ChannelPredictionEndEvent))
	case events.ChannelPredictionLockEvent:
		client.sendChannelPredictionLock(event.(events.ChannelPredictionLockEvent))
	case events.ChannelPredictionProgressEvent:
		client.sendChannelPredictionProgress(event.(events.ChannelPredictionProgressEvent))
	case events.ChannelRaidEvent:
		client.sendChannelRaid(event.(events.ChannelRaidEvent))
	case events.ChannelShieldModeBeginEvent:
		client.sendChannelShieldModeBegin(event.(events.ChannelShieldModeBeginEvent))
	case events.ChannelShieldModeEndEvent:
		client.sendChannelShieldModeEnd(event.(events.ChannelShieldModeEndEvent))
	case events.ChannelShoutoutCreateEvent:
		client.sendChannelShoutoutCreate(event.(events.ChannelShoutoutCreateEvent))
	case events.ChannelShoutoutReceiveEvent:
		client.sendChannelShoutoutReceive(event.(events.ChannelShoutoutReceiveEvent))
	case events.ChannelSubscribeEvent:
		client.sendChannelSubscribe(event.(events.ChannelSubscribeEvent))
	case events.ChannelSubscriptionEndEvent:
		client.sendChannelSubscriptionEnd(event.(events.ChannelSubscriptionEndEvent))
	case events.ChannelSubscriptionGiftedEvent:
		client.sendChannelSubscriptionGift(event.(events.ChannelSubscriptionGiftedEvent))
	case events.ChannelSubscriptionMessageEvent:
		client.sendChannelSubscriptionMessage(event.(events.ChannelSubscriptionMessageEvent))
	case events.ChannelUnbanEvent:
		client.sendChannelUnban(event.(events.ChannelUnbanEvent))
	case events.ChannelUpdateEvent:
		client.sendChannelUpdate(event.(events.ChannelUpdateEvent))
	case events.StreamOfflineEvent:
		client.sendStreamOffline(event.(events.StreamOfflineEvent))
	case events.StreamOnlineEvent:
		client.sendStreamOnline(event.(events.StreamOnlineEvent))
	case events.UserUpdateEvent:
		client.sendUserUpdate(event.(events.UserUpdateEvent))
	default:
		return
	}
}
