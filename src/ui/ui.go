package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width    int
	height   int
	keymap   keymap
	showHelp bool

	tabs       []string
	tabContent []string
	activeTab  int
}

func InitialModel() model {
	tabs := []string{"Player", "Library", "Settings"}
	tabContent := []string{"Player", "Library", "Settings"}
	return model{width: 300, height: 200, keymap: keys, tabs: tabs, tabContent: tabContent}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.tab):
			m.activeTab = (m.activeTab + 1) % len(m.tabs)
			return m, nil
		}
	}
	return m, nil
}

var (
	inactiveTabStyle = lipgloss.NewStyle().Bold(false).Foreground(lipgloss.Color("146")).Align(lipgloss.Left).PaddingBottom(1)
	activeTabStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("147")).Align(lipgloss.Left).PaddingBottom(1)
	docStyle         = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)

func (m model) View() string {
	var renderedTabs []string

	for i, t := range m.tabs {
		var style lipgloss.Style
		isActive := m.activeTab == i
		if isActive {
			style = activeTabStyle
		} else {
			style = inactiveTabStyle
		}

		renderedTabs = append(renderedTabs, style.Width(10).Render(t))
	}

	col := lipgloss.JoinVertical(lipgloss.Left, renderedTabs...)

	content := lipgloss.NewStyle().
		Padding(0, 2).
		Width(50).
		Height(m.height - 5).
		Render("Content for tab: " + m.tabs[m.activeTab])

	row := lipgloss.JoinHorizontal(lipgloss.Top, col, content)

	return docStyle.Render(row)
}
