package main

import (
	"bufio"
	"github.com/go-vgo/robotgo"
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

// type CommandAutomaton struct {
// 	I *CommandInterpretron
// 	commandStream chan string
// }

// func (ca *CommandAutomaton) executeCommands() error {
	
// }

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	commandinterpreter := NewInterpetron(allowed_keys)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		command := scanner.Text()

		if interpretedCommand, err := commandinterpreter.parseCommand(command); err == nil {
			robotgo.TypeString(string(interpretedCommand))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
