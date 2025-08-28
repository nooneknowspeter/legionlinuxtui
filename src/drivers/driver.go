package drivers

import (
	"fmt"
	"legionlinuxtui/src/helpers"
	"legionlinuxtui/src/system"
	"strconv"
)

type (
	DriverModuleFunction struct {
		File          string
		SysFSLocation string
		GetStatus     func(s *DriverModuleFunction) string
		UpdateStatus  func(s *DriverModuleFunction, value string)
		DriverModes   map[int]string
	}
)

var DEFAULTDRIVERMODES = map[int]string{
	0: "disabled",
	1: "enabled",
}

func (s *DriverModuleFunction) readValue() int {
	value, err := strconv.Atoi(helpers.ReadFile(s.SysFSLocation + s.File))
	if err != nil {
		panic(err)
	}

	return value
}

func (s *DriverModuleFunction) writeValue(value string) {
	helpers.WriteToFile(s.SysFSLocation+s.File, value)
}

// function implementation of iterator style pattern
// loop through modes in cycle; next
// currentMode value must be indexed at 0
func (s *DriverModuleFunction) toggleIterator(currentMode int, numberOfModes int) int {
	return ((1 + currentMode) % numberOfModes)
}

func (s *DriverModuleFunction) ToggleDriverMode() {
	currentMode := s.readValue()
	numberOfModes := len(s.DriverModes)
	var nextMode int

	if numberOfModes > 2 {
		nextMode = s.toggleIterator(currentMode-1, numberOfModes)
	} else {
		nextMode = s.toggleIterator(currentMode, numberOfModes)
	}

	s.writeValue(fmt.Sprintf("%v\n", nextMode))
	fmt.Printf("toggled %v: %v -> %v\n", s.File, s.DriverModes[currentMode], s.DriverModes[s.readValue()])
}

var (
	CameraPower DriverModuleFunction = DriverModuleFunction{
		File:          "camera_power",
		SysFSLocation: system.IDEASYSTEMDRIVERPATH,
		GetStatus: func(s *DriverModuleFunction) string {
			return s.DriverModes[s.readValue()]
		},
		DriverModes: DEFAULTDRIVERMODES,
	}

	ConservationMode DriverModuleFunction = DriverModuleFunction{
		File:          "conservation_mode",
		SysFSLocation: system.IDEASYSTEMDRIVERPATH,
		GetStatus: func(s *DriverModuleFunction) string {
			return s.DriverModes[s.readValue()]
		},
		DriverModes: DEFAULTDRIVERMODES,
	}
)
