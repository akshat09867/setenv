package bubble

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Screen int

const (
	WelcomeScreen Screen = iota
	ToolsScreen
)

type Mo struct {
	CurrentScreen Screen
	Choices       []string
	Cursor        int
	Selected      map[int]struct{}
}

func InitialModel() Mo {
	return Mo{
		CurrentScreen: WelcomeScreen,
		Choices:       []string{"Git", "Docker", "SSH"},
		Selected:      make(map[int]struct{}),
	}
}

func (m Mo) Init() tea.Cmd {
	return nil
}

func (m Mo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.CurrentScreen {
		case WelcomeScreen:
			switch msg.String() {
			case "enter", " ":
				m.CurrentScreen = ToolsScreen
				return m, nil
			case "q", "ctrl+c":
				return m, tea.Quit
			}

		case ToolsScreen:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "up", "k":
				if m.Cursor > 0 {
					m.Cursor--
				}
			case "down", "j":
				if m.Cursor < len(m.Choices)-1 {
					m.Cursor++
				}
			case "enter", " ":
				_, ok := m.Selected[m.Cursor]
				if ok {
					delete(m.Selected, m.Cursor)
				} else {
					m.Selected[m.Cursor] = struct{}{}
				}
			}
		}
	}
	return m, nil
}
