package main

import "github.com/charmbracelet/bubbles/key"

var keys = keymap{
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
