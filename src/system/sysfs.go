package system

import (
	"fmt"
	"legionlinuxtui/src/helpers"
	"strconv"
)

type (
	// DriverModuleFunction -> available functions in the legionlinux drivers
	DriverModuleFunction struct {
		File          string
		SysFSLocation string
		DriverStates  map[int]string
	}
)

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

// ToggleDriverState -> cycle through driver states
// loop through in cycle; next
func (s *DriverModuleFunction) ToggleDriverState() {
	currentState := s.readValue()
	numberOfStates := len(s.DriverStates)
	nextState := ((currentState + 1) % numberOfStates)

	// custom power mode
	if numberOfStates > 2 && currentState == 3 {
		nextState = 255
	}

	if numberOfStates > 2 && nextState == 0 {
		nextState++
	}

	s.writeValue(fmt.Sprintf("%v\n", nextState))
	fmt.Printf("toggled %v: %v -> %v\n", s.File, s.DriverStates[currentState], s.DriverStates[s.readValue()])
}

// GetDriverFunctionState -> get driver function current state from sysfs
func (s *DriverModuleFunction) GetDriverFunctionState() string {
	return s.DriverStates[s.readValue()]
}

// UpdateDriverFunction -> update driver function by value
func (s *DriverModuleFunction) UpdateDriverFunction(value string) {
	s.writeValue(value)
}

var (
	// DEFAULTDRIVERSTATES -> default driver function states; on and off
	DEFAULTDRIVERSTATES = map[int]string{
		0: "disabled",
		1: "enabled",
	}

	// ConservationMode ->
	ConservationMode = DriverModuleFunction{
		File:          "conservation_mode",
		SysFSLocation: IDEASYSTEMDRIVERPATH,
		DriverStates:  DEFAULTDRIVERSTATES,
	}

	// FNLock ->
	FNLock = DriverModuleFunction{
		File:          "fn_lock",
		SysFSLocation: IDEASYSTEMDRIVERPATH,
		DriverStates:  DEFAULTDRIVERSTATES,
	}

	// HybridMode -> gsync
	HybridMode = DriverModuleFunction{
		File:          "gsync",
		SysFSLocation: LEGIONSYSTEMDRIVERPATH,
		DriverStates:  DEFAULTDRIVERSTATES,
	}

	// LockFanController ->
	LockFanController = DriverModuleFunction{
		File:          "lockfancontroller",
		SysFSLocation: LEGIONSYSTEMDRIVERPATH,
		DriverStates:  DEFAULTDRIVERSTATES,
	}

	// MaxFanSpeed -> only works when power mode is in custom
	MaxFanSpeed = DriverModuleFunction{
		File:          "fan_fullspeed",
		SysFSLocation: LEGIONSYSTEMDRIVERPATH,
		DriverStates:  DEFAULTDRIVERSTATES,
	}

	// OverDrive ->
	OverDrive = DriverModuleFunction{
		File:          "overdrive",
		SysFSLocation: LEGIONSYSTEMDRIVERPATH,
		DriverStates:  DEFAULTDRIVERSTATES,
	}

	// PowerMode ->
	PowerMode = DriverModuleFunction{
		File:          "powermode",
		SysFSLocation: LEGIONSYSTEMDRIVERPATH,
		DriverStates: map[int]string{
			1:   "quiet",
			2:   "balanced",
			3:   "performance",
			255: "custom",
		},
	}

	// RapidCharge ->
	RapidCharge = DriverModuleFunction{
		File:          "rapidcharge",
		SysFSLocation: LEGIONSYSTEMDRIVERPATH,
		DriverStates:  DEFAULTDRIVERSTATES,
	}

	// TouchPad ->
	TouchPad = DriverModuleFunction{
		File:          "touchpad",
		SysFSLocation: LEGIONSYSTEMDRIVERPATH,
		DriverStates:  DEFAULTDRIVERSTATES,
	}

	// USBCharging ->
	USBCharging = DriverModuleFunction{
		File:          "usb_charging",
		SysFSLocation: IDEASYSTEMDRIVERPATH,
		DriverStates:  DEFAULTDRIVERSTATES,
	}

	// WinKey ->
	WinKey = DriverModuleFunction{
		File:          "winkey",
		SysFSLocation: LEGIONSYSTEMDRIVERPATH,
		DriverStates:  DEFAULTDRIVERSTATES,
	}
)
