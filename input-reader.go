package main

import (
	"bufio"
	"errors"
	"github.com/go-vgo/robotgo"
	"log"
	"os"
)

type Command string

type CommandInterpreter interface {
	parseCommand(string) (Command, error)
}

type CommandInterpretron struct{}

func (c *CommandInterpretron) parseCommand(chat string) (Command, error) {
	value, allowed := allowed_keys[chat]

	if allowed {
		return value, nil
	} else {
		return "", errors.New("Command not recognized.")
	}
}

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
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	commandinterpreter := new(CommandInterpretron)

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
