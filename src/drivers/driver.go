package drivers

import (
	"legionlinuxtui/src/helpers"
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
