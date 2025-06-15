package ui

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
	tab      key.Binding
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

var keys = keymap{
	tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "change tab"),
	),
	play: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "play"),
	),
	pause: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "pause"),
	),
	up: key.NewBinding(
		key.WithKeys("k"),
		key.WithHelp("k", "up"),
	),
	down: key.NewBinding(
		key.WithKeys("j"),
		key.WithHelp("j", "down"),
	),
	quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	forward: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("→/l", "seek +10s"),
	),
	backward: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("←/h", "seek -10s"),
	),
	next: key.NewBinding(
		key.WithKeys("]"),
		key.WithHelp("]", "next chapter"),
	),
	prev: key.NewBinding(
		key.WithKeys("["),
		key.WithHelp("[", "prev chapter"),
	),
	help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
}
