package main

import "fmt"

var allowed_keys = map[string]Command{
	"Left":  "left",
	"Right": "right",
	"Up":    "up",
	"Down":  "down",
	"A":     "z",
	"B":     "x",
	"Start": "enter",
}

func main() {
	interpreter := NewInterpetron(allowed_keys)
	kommander := newAutomaton(interpreter)

	messageCallback := func(message string) {
		err := kommander.executeCommand(message)
		fmt.Println(err)
	}

	getWindow("OpenEmu")

	ConnectToTwitch("wow", "test", messageCallback)
}
