package main

import (
	"bufio" 
	"log"
	"os"
)

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
	commander := newAutomaton(interpreter)

	commander.commandRoutine()
}
