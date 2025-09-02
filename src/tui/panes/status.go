package panes

import (
	"fmt"
	"legionlinuxtui/src/system"
	"legionlinuxtui/src/tui/components"
	"legionlinuxtui/src/tui/styles"

	"github.com/charmbracelet/lipgloss"
)

func StatusPane(terminalWidth int, terminalHeight int) string {
	systemInformation := components.StatusComponent(
		"System Information 󰌢 ",
		fmt.Sprintf("Host Name | %s", system.System.Name),
		fmt.Sprintf("Model | %s", system.System.Family),
		fmt.Sprintf("BIOS Version | %s", system.System.BiosVersion))

	cpuInformation := components.StatusComponent(
		"CPU  ",
		system.CPU.Name(),
		components.StreamlineTemperatureChart(system.CPU.Temperature()))

	gpuInformation := components.StatusComponent(
		"GPU 󰢮 ",
		system.GPU.Name(),
		components.StreamlineTemperatureChart(system.GPU.Temperature()))

	fanInformation := components.StatusComponent(
		"Fans 󰈐 ",
		fmt.Sprintf("Fan 1 | %vRPM",
			system.Fan1.CurrentSpeed()),
		fmt.Sprintf("Fan 2 | %vRPM",
			system.Fan2.CurrentSpeed()))

	batteryInformation := components.StatusComponent(
		"Battery 󰁹 ",
		components.BatteryInformationComponent(system.Battery.Capacity(), system.Battery.ChargingStatus()),
		fmt.Sprintf("Voltage | %sV", system.Battery.Voltage()),
		fmt.Sprintf("Cycle Count | %s", system.Battery.CycleCount()))

	return styles.Border.Width(terminalWidth).Render(lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.PlaceHorizontal(terminalWidth/5, lipgloss.Right, systemInformation),
		lipgloss.PlaceHorizontal(terminalWidth/5, lipgloss.Right, cpuInformation),
		lipgloss.PlaceHorizontal(terminalWidth/5, lipgloss.Right, gpuInformation),
		lipgloss.PlaceHorizontal(terminalWidth/5, lipgloss.Right, fanInformation),
		lipgloss.PlaceHorizontal(terminalWidth/5, lipgloss.Right, batteryInformation)))
}
