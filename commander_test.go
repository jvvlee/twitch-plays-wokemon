package main

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var mockError = errors.New("Unexpected chat")

type MockCommandInterpretron struct {
	expectedChat  string
	resultCommand Command
	resultError   error
}

func (c *MockCommandInterpretron) parseCommand(chat string) (Command, error) {
	if c.expectedChat == chat {
		return c.resultCommand, c.resultError
	}

	return "", mockError
}

func TestCommander(t *testing.T) {
	Convey("constructor returns a Commander", t, func() {
		mockInterpretor := new(MockCommandInterpretron)
		sut := newAutomaton(mockInterpretor)

		So(sut, ShouldImplement, (*Commander)(nil))
	})

	Convey("Passes errors from interpretron", t, func() {
		testError := errors.New("some error")
		mockInterpretor := MockCommandInterpretron{
			expectedChat:  "foo",
			resultCommand: "",
			resultError:   testError,
		}

		sut := newAutomaton(&mockInterpretor)
		err := sut.executeCommand("foo")
		So(err, ShouldEqual, testError)
	})
}
