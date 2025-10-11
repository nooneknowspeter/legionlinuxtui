package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	Heading               = lipgloss.NewStyle().Bold(true)
	Padding           int = 4
	Border                = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).Padding(Padding)
	TitleStyle            = lipgloss.NewStyle().MarginLeft(2).Bold(true)
	ItemStyle             = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle     = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#666666"))
	PaginationStyle       = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle             = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	QuitTextStyle         = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)
