package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Heading     = lipgloss.NewStyle().Bold(true)
	Padding int = 1
	Border      = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).Padding(Padding)
)
