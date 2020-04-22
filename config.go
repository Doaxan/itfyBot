package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Configuration struct {
	TelegramApiToken string
	TelegramApiDebug bool
	RssUpdateTimeSec time.Duration
	SendTimeoutMsec  time.Duration
}

var config Configuration

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
