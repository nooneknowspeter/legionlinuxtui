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

func TestGPUInformation(t *testing.T) {
	var (
		gpuName        string = system.GPU.Name()
		gpuTemperature string = system.GPU.Temperature()
	)

	if gpuName == "" || gpuTemperature == "" {
		t.Errorf("%v retrieve all gpu information ", Fail)
		return
	}

	fmt.Printf("GPU Model Name: %s\n", gpuName)
	fmt.Printf("GPU Temperature: %s°C\n", gpuTemperature)

	fmt.Printf("%v retrieved all gpu information ", Success)
}

func TestSystemInformation(t *testing.T) {
	var (
		hostName     string = system.System.Name
		biosVersion  string = system.System.BiosVersion
		family       string = system.System.Family
		systemVendor string = system.System.Vendor
	)

	if hostName == "" || biosVersion == "" || family == "" || systemVendor == "" {
		t.Errorf("%v retrieve all system information ", Fail)
		return
	}

	fmt.Printf("System Hostname : %s\n", hostName)
	fmt.Printf("System BIOS Version: %s\n", biosVersion)
	fmt.Printf("System Family: %s\n", family)
	fmt.Printf("System Vendor: %s\n", systemVendor)

	fmt.Printf("%v retrieved all system information ", Success)
}
