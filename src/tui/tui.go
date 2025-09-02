package tui

import (
	"log"
	"time"

	"legionlinuxtui/src/tui/panes"
	"legionlinuxtui/src/tui/styles"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	model struct {
		terminalWidth  int
		terminalHeight int
		windowTitle    string
		pollingRate    int // ms
	}

	tickMsg struct{}
)

func initModel() model {
	return model{
		pollingRate: 100,
	}
}

func ticker(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("legionlinuxtui"),
		ticker(time.Duration(m.pollingRate)*time.Millisecond),
	)
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.terminalHeight = msg.Height
		m.terminalWidth = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg:
		return m, ticker(time.Duration(m.pollingRate) * time.Millisecond)
	}

	return m, nil
}

func (m model) View() string {
	statusPane := panes.StatusPane(m.terminalWidth-(styles.Padding*2), m.terminalHeight)

	return statusPane
}

func Run() {
	p := tea.NewProgram(initModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
