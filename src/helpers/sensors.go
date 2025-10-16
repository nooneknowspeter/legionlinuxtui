package helpers

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func SensorInformation(sensorName string, sensorValueKey string) string {
	stdout, err := exec.Command("sensors", "-j").Output()
	if err != nil {
		fmt.Printf("command 'sensors' not found, ensure lm_sensors is installed %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(stdout, &data); err != nil {
		panic(err)
	}

	cpuTemperature := fmt.Sprintf("%v", data["legion_hwmon-isa-0000"].(map[string]any)[sensorName].(map[string]any)[sensorValueKey].(float64))

	return string(cpuTemperature)
}
