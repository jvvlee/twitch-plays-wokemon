package main

import "fmt"

var allowed_keys = map[string]Command{
	"Left":  "Left",
	"Right": "Right",
	"Up":    "Up",
	"Down":  "Down",
	"A":     "Z",
	"B":     "X",
	"Start": "Enter",
}

func main() {
	interpreter := NewInterpetron(allowed_keys)
	kommander := newAutomaton(interpreter)

	messageCallback := func(message string) {
		kommander.executeCommand(message)
		fmt.Println(message)
	}

	ConnectToTwitch("wow", "disguisedtoasths", messageCallback)
}
