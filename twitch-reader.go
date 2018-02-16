package main

import(
  log "github.com/Sirupsen/logrus"
  twitchchat "github.com/dimorinny/twitch-chat-api"
  "fmt"
)

var (
  config *twitchchat.Configuration
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

func main() {
  config := twitchchat.NewConfiguration(
    "test",
    "oauth:ex9ki2l39r579fkzlk7ry4h3zdb4hy",
    "test",
  )

  twitch := twitchchat.NewChat(config)

  stop := make(chan struct{})
  defer close(stop)

  err := twitch.ConnectWithCallbacks(connect, disconnect, newMessage)

  if err != nil {
    log.Fatal(err)
  }

  <- stop
}
