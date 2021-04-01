package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jcuga/golongpoll"
	"github.com/sirupsen/logrus"
)

type message struct {
	Type        string   `json:"type"`
	PlayerName  string   `json:"player"`
	TitleText   string   `json:"title,omitempty"`
	MessageText string   `json:"message,omitempty"`
	Command     string   `json:"command,omitempty"`
	Args        []string `json:"args,omitempty"`
}

type mtdBot struct {
	Config   *yamlConf
	LongPoll *golongpoll.LongpollManager
	Discord  *discordgo.Session
}

var (
	log        = logrus.New()
	configFile = flag.String("config", "config.yaml", "Config file")
)

func main() {
	flag.Parse()
	Config := getConfig()
	BOT := initBot(Config)

	// Correct shutdown
	go BOT.shutdown(make(chan os.Signal, 1))

	// Discord handler
	BOT.Discord.AddHandler(BOT.fromDiscord)

	// Serve our event subscription web handler
	http.HandleFunc("/sub", BOT.LongPoll.SubscriptionHandler)
	http.HandleFunc("/pub", BOT.fromMT)

	log.Println("Http server started.. [OK]")
	log.Fatal(http.ListenAndServe(Config.Host, nil))
}

func (b *mtdBot) shutdown(c chan os.Signal) {
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-c
	b.Discord.Close()     // Cleanly close down the Discord session
	b.LongPoll.Shutdown() // Stops the internal goroutine that provides subscription behavior
	log.Fatal("Stopping bot!")
}

func initBot(c *yamlConf) *mtdBot {
	mtbot := new(mtdBot)
	mtbot.Config = c
	mtbot.LongPoll = initLongPoll(c)
	mtbot.Discord = initDiscord(c)
	return mtbot
}
