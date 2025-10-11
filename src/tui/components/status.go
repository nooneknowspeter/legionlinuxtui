package components

import (
	"github.com/charmbracelet/lipgloss"
)

func StatusComponent(title string, information ...string) string {
	title = lipgloss.NewStyle().Bold(true).Render(title)

	var componentContent []string = []string{}

	componentContent = append(componentContent, title)

	for _, v := range information {
		componentContent = append(componentContent, v)
	}

	return lipgloss.NewStyle().Render(lipgloss.JoinVertical(lipgloss.Top, componentContent...))
}
