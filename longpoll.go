package main

import (
	"encoding/json"
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
		log.Fatalf("init longpoll server: %v", err)
	}

	log.Infoln("Initialization longpoll... [OK]")
	return manager
}

// Messages from minetest chat
func (b *mtdBot) fromMT(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	dec := json.NewDecoder(r.Body)
	msg := message{}
	err := dec.Decode(&msg)
	if err != nil {
		log.Errorf("unmarshalling: %v", err)
		return
	}
	// sending message to the Discord channel
	log.Debugln("MT >: ", msg)
	b.toDiscord(msg, keys.Get("category"))
}

func (b *mtdBot) toMT(m message, chID string) {
	log.Debugln("> MT: ", m)
	b.LongPoll.Publish(chID, m)
}
