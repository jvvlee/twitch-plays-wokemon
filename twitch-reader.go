package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	twitchchat "github.com/dimorinny/twitch-chat-api"
)

var config *twitchchat.Configuration = twitchchat.NewConfiguration(
	"wow",
	os.Getenv("TWITCH_OAUTH"),
	"swiftor",
)

func disconnect() {
	fmt.Println("Disconnected!")
}

func connect() {
	fmt.Println("Connected!")
}

func ConnectToTwitch(mh twitchchat.NewMessage) {
	twitch := twitchchat.NewChat(config)

	stop := make(chan struct{})
	defer close(stop)

	err := twitch.ConnectWithCallbacks(connect, disconnect, mh)

	if err != nil {
		log.Fatal(err)
	}

	<-stop
}
