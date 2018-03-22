package pokemon

import (
	"github.com/go-vgo/robotgo"
)

func GetWindow(emulatorName string) error {
	emulatorPidArray, err := robotgo.FindIds(emulatorName)

	if err != nil {
		return err
	}

	emulatorPid := emulatorPidArray[0]
	robotgo.ActivePID(emulatorPid)

	return nil
}
