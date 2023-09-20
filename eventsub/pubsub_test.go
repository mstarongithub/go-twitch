package eventsub

import (
	"os"
	"testing"
)

// TODO: Update tests

const chanID = 44322889

func TestShardedConnection(t *testing.T) {
	client := New()
	client.SetMaxShards(0)
	client.SetMaxTopicsPerShard(0)
	topic := Topic{
		Name:      "Heck if I know",
		Version:   1,
		ChannelID: chanID,
	}
	if err := client.Listen(topic); err != nil {
		t.Fatal(err)
	}
	t.Logf("listened to %d topics on %d shards", client.GetNumTopics(), client.GetNumShards())
	if err := client.Unlisten(topic); err != nil {
		t.Fatal(err)
	}
	client.Close()
}

func TestSingleConnection(t *testing.T) {
	conn := &Conn{}
	conn.SetMaxTopics(0)
	if err := conn.Connect(); err != nil {
		t.Fatal(err)
	}
	topic := Topic{
		Name:      "Heck if I know",
		Version:   1,
		ChannelID: chanID,
	}
	if err := conn.Listen(topic); err != nil {
		t.Fatal(err)
	}
	if err := conn.Reconnect(); err != nil {
		t.Fatal(err)
	}
	if err := conn.Unlisten(topic); err != nil {
		t.Fatal(err)
	}
	conn.Close()
}

func TestAuthenticatedConnection(t *testing.T) {
	token := os.Getenv("TWITCH_TOKEN")
	if len(token) < 1 {
		t.Skipf("TWITCH_TOKEN is not set. Skipping...")
	}
	client := New()
	topic := Topic{
		Name:      "Heck if I know",
		Version:   1,
		ChannelID: chanID,
	}
	if err := client.ListenWithAuth(token, topic); err != nil {
		t.Fatal(err)
	}
	if err := client.Unlisten(topic); err != nil {
		t.Fatal(err)
	}
	client.Close()
}
