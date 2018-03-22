package main

import (
	"github.com/go-vgo/robotgo"
)

func getWindow(emulatorName string) error {
	emulatorPidArray, err := robotgo.FindIds(emulatorName)

	if err != nil {
		return err
	}

	emulatorPid := emulatorPidArray[0]
	robotgo.ActivePID(emulatorPid)

	return nil
}
