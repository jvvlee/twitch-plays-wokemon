package main

import (
  "github.com/go-vgo/robotgo"
  "os"

)

type CommandAutomaton struct {
  interpreter *CommandInterpretron
}

func newAutomaton(interpetron *CommandInterpretron) *CommandAutomaton {
  automaton := new(CommandAutomaton)
  automaton.interpreter = interpetron
  return automaton
}

func (c *CommandAutomaton) executeCommand(kommand string) error {
  if interpretedCommand, err := c.interpreter.parseCommand(kommand); err == nil {
    robotgo.TypeString(string(interpretedCommand))
  }

  return err
}

func (c *CommandAutomaton) commandRoutine() error {
  file, err := os.Open("file.txt")
  if err != nil {
    log.Fatal(err)
  }

  //Eventually I will want to start reading from a channel.

  commandinterpreter := NewInterpetron(allowed_keys)

  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    command := scanner.Text()
    c.executeCommand(command)
  }

  if err := scanner.Err(); err != nil {
    return err
  }
}
