package bubble

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m Mo) View() tea.View {
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

	return tea.NewView(s)
}
