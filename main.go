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
	kommander := newAutomaton(interpreter, "file.txt")

	kommander.commandRoutine()

	fmt.Println("past kommander")
	th := new(twitchPrinter)
	ConnectToTwitch(th)
}
