package tests

import (
	"fmt"
	"legionlinuxtui/src/system"
	"testing"
)

func TestBatteryInformation(t *testing.T) {
	var (
		batteryName           string = system.Battery.Name()
		batteryCapacity       string = system.Battery.Capacity()
		batteryVoltage        string = system.Battery.Voltage()
		batteryChargingStatus string = system.Battery.ChargingStatus()
		batteryCycleCount     string = system.Battery.CycleCount()
	)

	if batteryName == "" || batteryCapacity == "" || batteryVoltage == "" || batteryChargingStatus == "" || batteryCycleCount == "" {
		t.Errorf("%v retrieve all battery information ", Fail)
		return
	}

	fmt.Printf("Battery Name: %s\n", batteryName)
	fmt.Printf("Battery Capacity 󰁿: %s%%\n", batteryCapacity)
	fmt.Printf("Battery Voltage 󱐋: %sV\n", batteryVoltage)
	fmt.Printf("Battery Charging Status: %s\n", batteryChargingStatus)
	fmt.Printf("Battery Cycle Count: %s\n", batteryCycleCount)

	fmt.Printf("%v retrieved all battery information ", Success)
}

func TestCPUInformation(t *testing.T) {
	var (
		cpuName        string = system.CPU.Name()
		cpuTemperature string = system.CPU.Temperature()
	)

	if cpuName == "" || cpuTemperature == "" {
		t.Errorf("%v retrieve all cpu information ", Fail)
		return
	}

	fmt.Printf("CPU Model Name : %s\n", cpuName)
	fmt.Printf("CPU Temperature : %s°C\n", cpuTemperature)

	fmt.Printf("%v retrieved all cpu information ", Success)
}
