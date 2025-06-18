package ui

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

type SettingsTab struct {
	currentDir   string
	entries      []fs.DirEntry
	cursor       int
	scrollOffset int
	viewHeight   int
	selected     []string
}

func NewSettingsTab(initialSelected []string, viewHeight int) SettingsTab {
	startDir, _ := os.UserHomeDir()
	allEntries, _ := os.ReadDir(startDir)

	var visibleEntries []fs.DirEntry
	for _, entry := range allEntries {
		if name := entry.Name(); len(name) > 0 && name[0] != '.' {
			visibleEntries = append(visibleEntries, entry)
		}
	}

	return SettingsTab{
		currentDir: startDir,
		entries:    visibleEntries,
		selected:   initialSelected,
		viewHeight: viewHeight,
	}
}

func (s SettingsTab) Init() tea.Cmd { return nil }
func (s SettingsTab) Name() string  { return "tab" }
func (s SettingsTab) Update(msg tea.Msg) (Tab, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		dirs := s.getDirs()

		switch msg.String() {
		case "k":
			if s.cursor > 0 {
				s.cursor--
				if s.cursor < s.scrollOffset {
					s.scrollOffset--
				}
			}
		case "j":
			if s.cursor < len(dirs)-1 {
				s.cursor++
				if s.cursor >= s.scrollOffset+s.viewHeight {
					s.scrollOffset++
				}
			}
			return s, nil
		case "h":
			s.changeDir(filepath.Dir(s.currentDir))
		case "l":
			if s.cursor < len(dirs) && dirs[s.cursor].IsDir() {
				s.changeDir(filepath.Join(s.currentDir, dirs[s.cursor].Name()))
			}
			return s, nil
		case "enter":
			if s.cursor < len(dirs) {
				path := filepath.Join(s.currentDir, dirs[s.cursor].Name())
				if !slices.Contains(s.selected, path) {
					s.selected = append(s.selected, path)
					return s, AddPathToConfig(path)
				}
			}
		}
	}
	return s, nil
}

func (s SettingsTab) View() string {
	var left strings.Builder
	var right strings.Builder

	left.WriteString(fmt.Sprintf("Browsing: %s\n\n", s.currentDir))
	right.WriteString("Selected Audiobook Folders:\n\n")

	dirs := s.getDirs()
	end := min(s.scrollOffset+s.viewHeight, len(dirs))

	for i := s.scrollOffset; i < end; i++ {
		cursor := "  "
		if i == s.cursor {
			cursor = "➤ "
		}
		left.WriteString(fmt.Sprintf("%s📁 %s\n", cursor, dirs[i].Name()))
	}

	if len(dirs) > s.viewHeight {
		left.WriteString(fmt.Sprintf("\nShowing %d-%d of %d\n", s.scrollOffset+1, end, len(dirs)))
	}

	for i := range min(s.viewHeight, len(s.selected)) {
		right.WriteString(fmt.Sprintf("✓ %s\n", s.selected[i]))
	}

	leftLines := strings.Split(left.String(), "\n")
	rightLines := strings.Split(right.String(), "\n")

	var final strings.Builder
	for i := range max(len(leftLines), len(rightLines)) {
		leftLine := ""
		rightLine := ""

		if i < len(leftLines) {
			leftLine = leftLines[i]
		}
		if i < len(rightLines) {
			rightLine = rightLines[i]
		}

		final.WriteString(fmt.Sprintf("%-40s  %s\n", leftLine, rightLine))
	}

	final.WriteString("\n↑/↓: Move  →: Enter folder  ←: Up  Enter: Select folder")
	return final.String()
}

func (s *SettingsTab) changeDir(path string) {
	allEntries, err := os.ReadDir(path)
	if err != nil {
		return
	}

	var visibleEntries []fs.DirEntry
	for _, entry := range allEntries {
		if name := entry.Name(); len(name) > 0 && name[0] != '.' {
			visibleEntries = append(visibleEntries, entry)
		}
	}
	s.currentDir = path
	s.entries = visibleEntries
	s.cursor = 0
}

func (s SettingsTab) Title() string {
	return "Settings"
}

func (s SettingsTab) getDirs() []fs.DirEntry {
	var dirs []fs.DirEntry
	for _, entry := range s.entries {
		if entry.IsDir() {
			dirs = append(dirs, entry)
		}
	}
	return dirs
}
