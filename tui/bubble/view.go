package bubble

import (
	"fmt"
)

func (m Mo) View() string {
	switch m.CurrentScreen {
	case WelcomeScreen:
		s := "WELCOME TO SETENV\n\n"
		s += "Press Enter to continue\n\n"
		s += "Press q to quit.\n"
		return s
	case roleScreen:
		s := "Select your role\n\n"
		for i, choice := range m.RoleChoices {
			cursor := ""
			if m.Cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
		s += "\nPress q to quit.\n"
		return s

	case ToolsScreen:
		s := "Select tools to check-\n\n"
		for i, choice := range m.Choices {
			cursor := " "
			if m.Cursor == i {
				cursor = ">"
			}
			checked := " "
			if _, ok := m.Selected[i]; ok {
				checked = "x"
			}
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
		}
		s += "\nPress q to quit.\n"
		return s
	}
	return ""
}
