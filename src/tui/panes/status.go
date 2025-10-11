package panes

import (
	"fmt"
	"legionlinuxtui/src/system"
	"legionlinuxtui/src/tui/components"
	"legionlinuxtui/src/tui/styles"

	"github.com/charmbracelet/lipgloss"
)

func StatusPane() string {
	systemInformation := components.StatusComponent(
		"System Information 󰌢 ",
		fmt.Sprintf("Host Name | %s", system.System.Name),
		fmt.Sprintf("Model | %s", system.System.Family),
		fmt.Sprintf("BIOS Version | %s", system.System.BiosVersion),
		fmt.Sprintf("CPU | %s %s", system.CPU.Name(), components.StreamlineTemperatureChart(system.CPU.Temperature())),
		fmt.Sprintf("GPU | %s %s", system.GPU.Name(), components.StreamlineTemperatureChart(system.GPU.Temperature())),
		fmt.Sprintf("Fan 1 | %vRPM", system.Fan1.CurrentSpeed()),
		fmt.Sprintf("Fan 2 | %vRPM", system.Fan2.CurrentSpeed()),
		fmt.Sprintf("Battery Capacity | %s", components.BatteryInformationComponent(system.Battery.Capacity(), system.Battery.ChargingStatus())),
		fmt.Sprintf("Battery Voltage | %sV", system.Battery.Voltage()),
		fmt.Sprintf("Battery Cycle Count | %s", system.Battery.CycleCount()),
	)

	return lipgloss.NewStyle().Padding(0, styles.Padding).Render(
		systemInformation,
	)
}
