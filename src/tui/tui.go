package tui

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"legionlinuxtui/src/system"
	"legionlinuxtui/src/tui/panes"
	"legionlinuxtui/src/tui/styles"
)

var (
	title           = "legionlinuxtui"
	driverFunctions = map[string]*system.DriverModuleFunction{
		"Conservation Mode":   &system.ConservationMode,
		"FNLock":              &system.FNLock,
		"Hybrid Mode":         &system.HybridMode,
		"Lock Fan Controller": &system.LockFanController,
		"Max Fan Speed (PowerMode must be set to 'custom')": &system.MaxFanSpeed,
		"OverDrive":      &system.OverDrive,
		"PowerMode":      &system.PowerMode,
		"Rapid Charging": &system.RapidCharge,
		"TouchPad":       &system.TouchPad,
		"USB Charging":   &system.USBCharging,
		"WinKey":         &system.WinKey,
	}
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s -> %s", i, driverFunctions[string(i)].GetDriverFunctionState())

	fn := styles.ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return styles.SelectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	_, err := fmt.Fprint(w, fn(str))
	if err != nil {
		fmt.Println("failed to print item")
	}
}

type (
	model struct {
		list           list.Model
		choice         string
		quitting       bool
		pollingRate    int // ms
		terminalWidth  int
		terminalHeight int
	}

	tickMsg struct{}
)

func ticker(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle(title),
		ticker(time.Duration(m.pollingRate)*time.Millisecond),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminalHeight = msg.Height
		m.terminalWidth = msg.Width
		m.list.SetHeight(m.terminalHeight - (lipgloss.Height(styles.TitleStyle.Render(title)) + lipgloss.Height("\n") + lipgloss.Height(panes.StatusPane())))
		m.list.SetWidth(m.terminalWidth)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
				driverFunctions[m.choice].ToggleDriverState()
				_ = tea.ClearScreen()
			}

			// return m, tea.Quit
		}

	case tickMsg:
		return m, ticker(time.Duration(m.pollingRate) * time.Millisecond)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	// if m.choice != "" {
	// 	return styles.QuitTextStyle.Render(fmt.Sprintf("Enabling/Disabling -> %s", m.choice))
	// }

	return lipgloss.JoinVertical(lipgloss.Left, styles.TitleStyle.Render(title), "\n", panes.StatusPane(), m.list.View())
}

func Run() {
	items := []list.Item{}

	for k := range driverFunctions {
		items = append(items, item(k))
	}

	const listHeight = 20
	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = ""
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(false)
	l.Styles.Title = styles.TitleStyle
	l.Styles.PaginationStyle = styles.PaginationStyle
	l.Styles.HelpStyle = styles.HelpStyle

	m := model{list: l, pollingRate: 100}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
