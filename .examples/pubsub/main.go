package main

import (
	"fmt"
	"github.com/Adeithe/go-twitch/eventsub"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Adeithe/go-twitch"
)

var mgr *eventsub.Client

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	mgr = twitch.PubSub()

	mgr.OnShardConnect(func(shard int) {
		fmt.Printf("Shard #%d connected!\n", shard)
	})

	mgr.OnShardReconnect(func(shard int) {
		fmt.Printf("Shard #%d reconnected!\n", shard)
	})

	mgr.OnShardMessage(func(shard int, topic string, data []byte) {
		fmt.Printf("Shard #%d > %s %s\n", shard, topic, strings.TrimSpace(string(data)))
	})

	mgr.OnShardLatencyUpdate(func(shard int, latency time.Duration) {
		fmt.Printf("Shard #%d has %.3fs ping!\n", shard, latency.Seconds())
	})

	mgr.OnShardDisconnect(func(shard int) {
		fmt.Printf("Shard #%d disconnected!\n", shard)
	})

	channelID := 44322889
	printErr(mgr.Listen("radio-events-pubsub", channelID))
	printErr(mgr.Listen("polls", channelID))
	printErr(mgr.Listen("hype-train-events-pubsub", channelID))
	printErr(mgr.Listen("video-playback-by-id", channelID))
	printErr(mgr.Listen("stream-chat-room-pubsub", channelID))
	printErr(mgr.Listen("community-points-channel-pubsub", channelID))
	printErr(mgr.Listen("pv-watch-party-events", channelID))
	printErr(mgr.Listen("extension-control", channelID))

	fmt.Printf("Started listening to %d topics on %d shards!\n", mgr.GetNumTopics(), mgr.GetNumShards())

	<-sc
	fmt.Println("Stopping...")
	mgr.Close()
}

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
