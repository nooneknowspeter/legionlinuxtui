package components

import (
	"github.com/charmbracelet/lipgloss"
)

func StatusComponent(title string, information ...string) string {
	title = lipgloss.NewStyle().Bold(true).Render(title)

	var componentContent = []string{}

	componentContent = append(componentContent, title)

	componentContent = append(componentContent, information...)

	return lipgloss.NewStyle().Render(lipgloss.JoinVertical(lipgloss.Top, componentContent...))
}
