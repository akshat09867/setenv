package bubble

import (
	tea "charm.land/bubbletea/v2"
)

type Mo struct {
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func InitialModel() Mo {
	return Mo{
		Choices: []string{"Git", "Docker", "ssh"},

		Selected: make(map[int]struct{}),
	}
}

func (m Mo) Init() tea.Cmd {
	return nil
}

func (m Mo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:

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

		case "enter", "space":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}
		}
	}

	return m, nil
}
