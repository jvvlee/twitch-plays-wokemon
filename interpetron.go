package main

import (
  "errors"
)

type Command string

type CommandInterpreter interface {
  parseCommand(string) (Command, error)
}

type CommandInterpretron struct{
  dict map[string]Command
}

func NewInterpetron(newDictionary map[string]Command) *CommandInterpretron {
  interpetron := new(CommandInterpretron)
  interpetron.dict = newDictionary
  return interpetron
}

func (c *CommandInterpretron) parseCommand(chat string) (Command, error) {
  value, allowed := allowed_keys[chat]

  if allowed {
    return value, nil
  } else {
    return "", errors.New("Command not recognized.")
  }
}