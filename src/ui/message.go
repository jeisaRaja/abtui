package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AddPathToConfigMsg struct {
	Path string
}

func AddPathToConfig(path string) tea.Cmd {
	return func() tea.Msg {
		return AddPathToConfigMsg{Path: path}
	}
}
