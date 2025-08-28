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
		t.Errorf("  retrieve all battery information ")
		return
	}

	fmt.Printf("Battery Name: %s\n", batteryName)
	fmt.Printf("Battery Capacity 󰁿: %s%%\n", batteryCapacity)
	fmt.Printf("Battery Voltage 󱐋: %sV\n", batteryVoltage)
	fmt.Printf("Battery Charging Status: %s\n", batteryChargingStatus)
	fmt.Printf("Battery Cycle Count: %s\n", batteryCycleCount)
}

func TestCPUInformation(t *testing.T) {
	var (
		cpuName        string = system.CPU.Name()
		cpuTemperature string = system.CPU.Temperature()
	)

	if cpuName == "" || cpuTemperature == "" {
		t.Errorf("  retrieve all cpu information ")
		return
	}

	fmt.Printf("CPU Model Name : %s\n", cpuName)
	fmt.Printf("CPU Temperature : %s°C\n", cpuTemperature)
}

func TestGPUInformation(t *testing.T) {
	var (
		gpuName        string = system.GPU.Name()
		gpuTemperature string = system.GPU.Temperature()
	)

	if gpuName == "" || gpuTemperature == "" {
		t.Errorf("  retrieve all gpu information ")
		return
	}

	fmt.Printf("GPU Model Name: %s\n", gpuName)
	fmt.Printf("GPU Temperature: %s°C\n", gpuTemperature)
}

func TestSystemInformation(t *testing.T) {
	var (
		hostName     string = system.System.Name
		biosVersion  string = system.System.BiosVersion
		family       string = system.System.Family
		systemVendor string = system.System.Vendor
	)

	if hostName == "" || biosVersion == "" || family == "" || systemVendor == "" {
		t.Errorf("  retrieve all system information ")
		return
	}

	fmt.Printf("System Hostname : %s\n", hostName)
	fmt.Printf("System BIOS Version: %s\n", biosVersion)
	fmt.Printf("System Family: %s\n", family)
	fmt.Printf("System Vendor: %s\n", systemVendor)
}
