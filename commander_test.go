package main

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type MockCommandInterpretron struct{}

func (c *MockCommandInterpretron) parseCommand(chat string) (Command, error) {
	return "", errors.New("Some error")
}

func TestCommander(t *testing.T) {
	Convey("constructor returns a Commander", t, func() {
		mockInterpretor := new(MockCommandInterpretron)
		mockCommander := newAutomaton(mockInterpretor)

		So(mockCommander, ShouldImplement, (*Commander)(nil))
		//Wtf is this???
	})

	Convey("Passes errors from interpretron", t, func() {
		mockInterpretor := new(MockCommandInterpretron)
		mockCommander := newAutomaton(mockInterpretor)

		err := mockCommander.executeCommand("")

		So(err, ShouldBeError)
	})
}
