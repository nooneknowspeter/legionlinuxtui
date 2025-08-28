package system

import (
	"encoding/json"
	"fmt"
	"legionlinuxtui/src/helpers"
	"os/exec"
)

type (
	systemInformation struct {
		Name        string
		BiosVersion string
		Family      string
		Vendor      string
	}

	cpuSensorInformation struct {
		Name        func() string
		Temperature func() string
	}
)

var (
	System systemInformation = systemInformation{
		Name:        helpers.ReadFile("/etc/hostname"),
		BiosVersion: helpers.ReadFile(DEVICEINFORMATIONPATH + "bios_version"),
		Family:      helpers.ReadFile(DEVICEINFORMATIONPATH + "product_family"),
		Vendor:      helpers.ReadFile(DEVICEINFORMATIONPATH + "sys_vendor"),
	}

	CPU cpuSensorInformation = cpuSensorInformation{
		Name: func() string {
			stdout, err := exec.Command("lscpu", "-e=ModelName", "-J").Output()
			if err != nil {
				fmt.Printf("command lscpu not found %v", err)
			}

			var data map[string]any
			if err := json.Unmarshal(stdout, &data); err != nil {
				panic(err)
			}

			cpuModel := fmt.Sprintf("%v", data["cpus"].([]any)[0].(map[string]any)["modelname"].(string))

			return cpuModel
		},

		Temperature: func() string {
			file := helpers.ReadFile(CPUINFORMATIONPATH + "temp1_input")

			cpuTemperature := fmt.Sprintf("%v", string(file[0:2]))

			return cpuTemperature
		},
	}
)
