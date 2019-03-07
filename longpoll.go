package main

import (
	"encoding/json"
	"io/ioutil"
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[Error] reading the request body", err)
		return
	}

	msg := message{}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Println("[Error] unmarshalling", err)
		return
	}
	// sending message to the Discord
	b.toDiscord(msg)
}

func (b *mtdBot) toMT(m message) {
	b.LongPoll.Publish(b.Config.LongPoll.Category, m)
}
