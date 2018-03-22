package main

import (
	"fmt"

	"github.com/personal/pwitch_plays_tokemon/pokemon"
)

var allowed_keys = map[string]pokemon.Command{
	"Left":  "left",
	"Right": "right",
	"Up":    "up",
	"Down":  "down",
	"A":     "z",
	"B":     "x",
	"Start": "enter",
}

func main() {
	interpreter := pokemon.NewInterpetron(allowed_keys)
	kommander := pokemon.NewAutomaton(interpreter)

	messageCallback := func(message string) {
		err := kommander.ExecuteCommand(message)
		fmt.Println(err)
	}

	pokemon.GetWindow("OpenEmu")

	pokemon.ConnectToTwitch("wow", "test", messageCallback)
}
