package main

import (
	"github.com/go-vgo/robotgo"
)

type Commander interface {
	executeCommand(string) error
}

type CommandAutomaton struct {
	interpreter CommandInterpreter
}

func newAutomaton(interpetron CommandInterpreter) Commander {
	automaton := new(CommandAutomaton)
	automaton.interpreter = interpetron
	return automaton
}

func (c *CommandAutomaton) executeCommand(kommand string) error {
	interpretedCommand, err := c.interpreter.parseCommand(kommand)

	if err == nil {
		robotgo.KeyTap(string(interpretedCommand))
		return nil
	}

	return err
}
