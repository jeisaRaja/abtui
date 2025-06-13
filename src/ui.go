package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keymap struct {
	play     key.Binding
	pause    key.Binding
	up       key.Binding
	down     key.Binding
	quit     key.Binding
	forward  key.Binding
	backward key.Binding
	next     key.Binding
	prev     key.Binding
	help     key.Binding
}

type model struct {
	width    int
	height   int
	keymap   keymap
	showHelp bool
}

func InitialModel() model {
	return model{width: 300, height: 200, keymap: keys}
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
		}
	}
	return m, nil
}

var modeStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("147")).Width(40).Align(lipgloss.Center)

func (m model) View() string {
	content := modeStyle.Render("abtui")
	box := lipgloss.NewStyle().
		AlignHorizontal(lipgloss.Center).
		Width(m.width).
		BorderForeground(lipgloss.Color("63")).
		Render(content)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}
