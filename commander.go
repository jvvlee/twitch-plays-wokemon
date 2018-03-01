package main

import (
  "github.com/go-vgo/robotgo"
  "os"
  "log"
  "bufio"
)

type CommandAutomaton struct {
  interpreter *CommandInterpretron
  filename string
}

func newAutomaton(interpetron *CommandInterpretron, filename string) *CommandAutomaton {
  automaton := new(CommandAutomaton)
  automaton.interpreter = interpetron
  automaton.filename = filename
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

func (c *CommandAutomaton) commandRoutine() error {
  file, err := os.Open(c.filename)

  if err != nil {
    log.Fatal(err)
  }

  //Eventually I will want to start reading from a channel.
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    command := scanner.Text()
    c.executeCommand(command)
  }

  if err := scanner.Err(); err != nil {
    return err
  }

  return nil
}
