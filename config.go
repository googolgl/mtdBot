package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type yamlConf struct {
	Host    string `yaml:"host"`
	Discord struct {
		Token string `yaml:"token"`
		//ChannelID string   `yaml:"channel"`
		Commands []string `yaml:"commands"`
		Help     string   `yaml:"help"`
	} `yaml:"discord"`
	LongPoll struct {
		//Category string        `yaml:"category"`
		Conn *connLongPoll `yaml:"connect"`
	} `yaml:"longPoll"`
}

type connLongPoll struct {
	LoggingEnabled                 bool `yaml:"loggingEnabled"`
	MaxLongpollTimeoutSeconds      int  `yaml:"maxLongpollTimeoutSeconds"`
	MaxEventBufferSize             int  `yaml:"maxEventBufferSize"`
	EventTimeToLiveSeconds         int  `yaml:"eventTimeToLiveSeconds"`
	DeleteEventAfterFirstRetrieval bool `yaml:"deleteEventAfterFirstRetrieval"`
}

func getConfig() *yamlConf {
	c := yamlConf{}
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal("[Error] reading file", err)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatal("[Error] unmarshalling", err)
	}
	log.Println("Reading config... [OK]")
	return &c
}
