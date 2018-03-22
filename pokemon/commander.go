package pokemon

import (
	"time"

	"github.com/go-vgo/robotgo"
)

type Commander interface {
	ExecuteCommand(string) error
}

type CommandAutomaton struct {
	interpreter CommandInterpreter
}

func NewAutomaton(interpetron CommandInterpreter) Commander {
	automaton := new(CommandAutomaton)
	automaton.interpreter = interpetron
	return automaton
}

func (c *CommandAutomaton) ExecuteCommand(kommand string) error {
	interpretedCommand, err := c.interpreter.parseCommand(kommand)

	if err == nil {
		robotgo.KeyToggle(string(interpretedCommand), "down")
		time.Sleep(1000 * time.Millisecond)
		robotgo.KeyToggle(string(interpretedCommand), "up")
		return nil
	}

	return err
}
