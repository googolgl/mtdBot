package main

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type yamlConf struct {
	Host    string `yaml:"host"`
	Discord struct {
		Token    string   `yaml:"token"`
		Commands []string `yaml:"commands"`
		Help     string   `yaml:"help"`
	} `yaml:"discord"`
	LongPoll struct {
		Conn *connLongPoll `yaml:"connect"`
	} `yaml:"longPoll"`
	LogLevel *logrus.Level `yaml:"loglevel"`
}

type connLongPoll struct {
	LoggingEnabled                 bool `yaml:"loggingEnabled"`
	MaxLongpollTimeoutSeconds      int  `yaml:"maxLongpollTimeoutSeconds"`
	MaxEventBufferSize             int  `yaml:"maxEventBufferSize"`
	EventTimeToLiveSeconds         int  `yaml:"eventTimeToLiveSeconds"`
	DeleteEventAfterFirstRetrieval bool `yaml:"deleteEventAfterFirstRetrieval"`
}

func getConfig() *yamlConf {
	// config log
	log.Formatter.(*logrus.TextFormatter).DisableColors = true
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true

	c := yamlConf{}
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("reading file %v", err)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("unmarshalling %v", err)
	}
	log.Infoln("Reading config... [OK]")

	// set log level
	log.SetLevel(*c.LogLevel)

	return &c
}
