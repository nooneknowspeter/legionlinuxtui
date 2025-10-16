package system

import (
	"encoding/json"
	"fmt"
	"legionlinuxtui/src/helpers"
	"os/exec"
	"strings"
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

	gpuSensorInformation struct {
		Name        func() string
		Temperature func() string
	}

	batterySensorInformation struct {
		Name           func() string
		Capacity       func() string
		Voltage        func() string
		ChargingStatus func() string
		CycleCount     func() string
	}

	fanSensorInformation struct {
		CurrentSpeed func() string
	}
)

var (
	System = systemInformation{
		Name:        helpers.ReadFile("/etc/hostname"),
		BiosVersion: helpers.ReadFile(DEVICEINFORMATIONPATH + "bios_version"),
		Family:      helpers.ReadFile(DEVICEINFORMATIONPATH + "product_family"),
		Vendor:      helpers.ReadFile(DEVICEINFORMATIONPATH + "sys_vendor"),
	}

	// CPU -> cpu information from `lscpu` command
	CPU = cpuSensorInformation{
		Name: func() string {
			stdout, err := exec.Command("lscpu", "-e=ModelName", "-J").Output()
			if err != nil {
				fmt.Printf("command 'lscpu' not found %v", err)
			}

			var data map[string]any
			if err := json.Unmarshal(stdout, &data); err != nil {
				panic(err)
			}

			cpuModel := fmt.Sprintf("%v", data["cpus"].([]any)[0].(map[string]any)["modelname"].(string))

			return cpuModel
		},

		Temperature: func() string {
			return helpers.SensorInformation("CPU Temperature", "temp1_input")
		},
	}

	// GPU -> gpu information from `nvidia-smi` command
	GPU = gpuSensorInformation{
		Name: func() string {
			stdout, err := exec.Command("nvidia-smi", "--query-gpu=name", "--format=csv,noheader").Output()
			if err != nil {
				fmt.Printf("command 'nvidia-smi' not found: %v", err)
			}

			return strings.Split(string(stdout), "\n")[0]
		},

		Temperature: func() string {
			stdout, err := exec.Command("nvidia-smi", "--query-gpu=temperature.gpu", "--format=csv,noheader").Output()
			if err != nil {
				panic(err)
			}

			return strings.Split(string(stdout), "\n")[0]
		},
	}

	Battery = batterySensorInformation{
		Name: func() string {
			file := helpers.ReadFile(BATTERYINFORMATIONPATH + "model_name")

			return file
		},

		Capacity: func() string {
			file := helpers.ReadFile(BATTERYINFORMATIONPATH + "capacity")

			return file
		},

		Voltage: func() string {
			file := helpers.ReadFile(BATTERYINFORMATIONPATH + "voltage_now")

			voltage := fmt.Sprintf("%v.%v", string(file[0:2]), string(file[2]))

			return voltage
		},

		ChargingStatus: func() string {
			file := helpers.ReadFile(BATTERYINFORMATIONPATH + "status")

			return file
		},

		CycleCount: func() string {
			file := helpers.ReadFile(BATTERYINFORMATIONPATH + "cycle_count")

			return file
		},
	}

	Fan1 = fanSensorInformation{
		CurrentSpeed: func() string {
			return helpers.SensorInformation("Fan 1", "fan1_input")
		},
	}

	Fan2 = fanSensorInformation{
		CurrentSpeed: func() string {
			return helpers.SensorInformation("Fan 2", "fan2_input")
		},
	}
)
