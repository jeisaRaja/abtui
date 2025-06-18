package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Tab interface {
	Init() tea.Cmd
	Update(tea.Msg) (Tab, tea.Cmd)
	View() string
	Name() string
}
