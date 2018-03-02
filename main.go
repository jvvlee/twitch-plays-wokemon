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

	newMessage := func(message string) {
		fmt.Println(message)
	}

	fmt.Println("past kommander")
	ConnectToTwitch(newMessage)
}
