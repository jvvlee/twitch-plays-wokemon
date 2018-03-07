package main

import (
	"github.com/go-vgo/robotgo"
)

type Commander interface {
	executeCommand(string) error
}

type CommandAutomaton struct {
	interpreter *CommandInterpretron
}

func newAutomaton(interpetron *CommandInterpretron) Commander {
	automaton := new(CommandAutomaton)
	automaton.interpreter = interpetron
	return automaton
}

func (c *CommandAutomaton) executeCommand(kommand string) error {
	interpretedCommand, err := c.interpreter.parseCommand(kommand)

	if err == nil {
		robotgo.TypeString(string(interpretedCommand))
		return nil
	}

	return err
}
