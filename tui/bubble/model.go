package bubble

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Screen int

const (
	WelcomeScreen Screen = iota
	roleScreen
	ToolsScreen
)

type Mo struct {
	CurrentScreen Screen
	NextScreen    Screen
	RoleChoices   []string
	Choices       []string
	ChosenRole    int
	Cursor        int
	Selected      map[int]struct{}
}

func InitialModel() Mo {
	return Mo{
		CurrentScreen: WelcomeScreen,
		NextScreen:    roleScreen,
		ChosenRole:    -1,
		RoleChoices:   []string{"Developer", "Designer", "SysAdmin"},
		Choices:       []string{"Git", "Docker", "SSH"},
		Selected:      make(map[int]struct{}),
	}
}

func (m Mo) Init() tea.Cmd {
	return nil
}
func (m *Mo) applyRecommendations() {
	m.Selected = make(map[int]struct{})
	if m.ChosenRole < 0 || m.ChosenRole >= len(m.RoleChoices) {
		return
	}
	role := m.RoleChoices[m.ChosenRole]
	var recommended []string
	switch role {
	case "Developer":
		recommended = []string{"Git", "Docker", "SSH"}
	case "Designer":
		recommended = []string{"Git"}
	case "SysAdmin":
		recommended = []string{"Docker", "SSH"}
	default:
		recommended = []string{"Git", "Docker", "SSH"}
	}
	for _, rec := range recommended {
		for i, tool := range m.Choices {
			if tool == rec {
				m.Selected[i] = struct{}{}
				break
			}
		}
	}
}
func (m Mo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.CurrentScreen {
		case WelcomeScreen:
			switch msg.String() {
			case "enter", " ":
				m.CurrentScreen = roleScreen
				m.Cursor = 0
				return m, nil
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		case roleScreen:
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
				m.ChosenRole = m.Cursor
				m.applyRecommendations()
				m.CurrentScreen = ToolsScreen
				m.Cursor = 0
				return m, nil
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
