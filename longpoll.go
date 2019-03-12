package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jcuga/golongpoll"
)

func initLongPoll(c *yamlConf) *golongpoll.LongpollManager {
	manager, err := golongpoll.StartLongpoll(golongpoll.Options{
		LoggingEnabled:                 c.LongPoll.Conn.LoggingEnabled,
		MaxLongpollTimeoutSeconds:      c.LongPoll.Conn.MaxLongpollTimeoutSeconds,
		MaxEventBufferSize:             c.LongPoll.Conn.MaxEventBufferSize,
		EventTimeToLiveSeconds:         c.LongPoll.Conn.EventTimeToLiveSeconds,
		DeleteEventAfterFirstRetrieval: c.LongPoll.Conn.DeleteEventAfterFirstRetrieval,
	})
	if err != nil {
		log.Fatal("[Error] init longpoll server", err)
	}

	log.Println("Initialization longpoll... [OK]")
	return manager
}

// Messages from minetest chat
func (b *mtdBot) fromMT(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	//log.Println(keys.Get("token"), keys.Get("category"))

	dec := json.NewDecoder(r.Body)
	msg := message{}
	err := dec.Decode(&msg)
	if err != nil {
		log.Println("[Error] unmarshalling", err)
		return
	}
	// sending message to the Discord channel
	b.toDiscord(msg, keys.Get("category"))
}

func (b *mtdBot) toMT(m message, chID string) {
	b.LongPoll.Publish(chID, m)
}
