package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	twitchchat "github.com/dimorinny/twitch-chat-api"
	"os"
)

var config *twitchchat.Configuration = twitchchat.NewConfiguration(
	"testicles",
	os.Getenv("TWITCH_OAUTH"),
	"data_dave",
)

func disconnect() {
	fmt.Println("Disconnected!")
}

func connect() {
	fmt.Println("Connected!")
}

func newMessage(message string) {
	fmt.Println(message)
}

func startTwitch() {

	twitch := twitchchat.NewChat(config)

	stop := make(chan struct{})
	defer close(stop)

	err := twitch.ConnectWithCallbacks(connect, disconnect, newMessage)

	if err != nil {
		log.Fatal(err)
	}

	<-stop
}

type TwitchStream struct {
	twitchConfig *twitchchat.Configuration
	twitchChannel chan string
}

