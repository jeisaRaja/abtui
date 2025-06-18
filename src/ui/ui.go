package ui

import (
	"slices"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jeisaRaja/abtui/src/config"
)

type model struct {
	width    int
	height   int
	keymap   keymap
	showHelp bool

	tabs             []string
	tabContent       []string
	activeTabContent Tab
	activeTab        int

	config *config.Config
}

func InitialModel() model {
	tabs := []string{"Player", "Library", "Settings"}
	tabContent := []string{"Player", "Library", "Settings"}
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		panic("config not found")
	}
	return model{width: 300, height: 200, keymap: keys, tabs: tabs, tabContent: tabContent, config: &cfg}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case AddPathToConfigMsg:
		if !slices.Contains(m.config.Filepaths, msg.Path) {
			m.config.Filepaths = append(m.config.Filepaths, msg.Path)
		}
		cfg, err := config.SaveConfig("config.json", m.config)
		if err != nil {
			panic("config not found")
		}
		m.config = cfg
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.tab):
			m.activeTab = (m.activeTab + 1) % len(m.tabs)
			switch m.tabs[m.activeTab] {
			case "Settings":
				m.activeTabContent = NewSettingsTab(m.config.Filepaths, m.height-10)
			}

		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		}
	}

	if m.activeTabContent != nil {

		tab, cmd := m.activeTabContent.Update(msg)
		m.activeTabContent = tab

		return m, cmd
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

	var content string
	if m.activeTabContent != nil {
		content = m.activeTabContent.View()
	} else {
		content = "content is still empty"
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, col, content)

	return docStyle.Render(row)
}
