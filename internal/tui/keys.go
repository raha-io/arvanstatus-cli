package tui

import "charm.land/bubbles/v2/key"

type keyMap struct {
	Up      key.Binding
	Down    key.Binding
	Tab     key.Binding
	Enter   key.Binding
	Esc     key.Binding
	Refresh key.Binding
	Help    key.Binding
	Quit    key.Binding
}

func defaultKeys() keyMap {
	return keyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		Tab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch pane"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "detail"),
		),
		Esc: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "close"),
		),
		Refresh: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "refresh"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Tab, k.Enter, k.Refresh, k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Tab},
		{k.Enter, k.Esc, k.Refresh},
		{k.Help, k.Quit},
	}
}
