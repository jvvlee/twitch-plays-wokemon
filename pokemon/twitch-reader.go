package pokemon

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	twitchchat "github.com/dimorinny/twitch-chat-api"
)

func disconnect() {
	fmt.Println("Disconnected!")
}

func connect() {
	fmt.Println("Connected!")
}

func ConnectToTwitch(nickname string, chatroom string, mh twitchchat.NewMessage) {
	config := twitchchat.NewConfiguration(
		nickname,
		os.Getenv("TWITCH_OAUTH"),
		chatroom,
	)

	twitch := twitchchat.NewChat(config)

	stop := make(chan struct{})

	defer close(stop)

	err := twitch.ConnectWithCallbacks(connect, disconnect, mh)

	if err != nil {
		log.Fatal(err)
	}

	<-stop
}
