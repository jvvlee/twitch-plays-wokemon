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
	"feelsgoodman",
)

type twitchHandler interface {
	newMessage(string)
}

type twitchPrinter struct {
}

func disconnect() {
	fmt.Println("Disconnected!")
}

func connect() {
	fmt.Println("Connected!")
}

func (t *twitchPrinter) newMessage(message string) {
	fmt.Println(message)
}

func ConnectToTwitch(th twitchHandler) {

	twitch := twitchchat.NewChat(config)

	stop := make(chan struct{})
	defer close(stop)

	err := twitch.ConnectWithCallbacks(connect, disconnect, th.newMessage)

	if err != nil {
		log.Fatal(err)
	}

	<-stop
}
