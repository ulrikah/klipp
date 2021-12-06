package klipp

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	commands []string
	cursor   int
	klipp    Klipp
	content  string
}

func initialModel() model {
	return model{
		commands: []string{"List notes", "Paste from buffer", "Read note"},
		cursor:   0,
		// these probably shouldn't be in the model
		klipp:   Klipp{HomeDir: "/Users/ulrikah/.klipp"},
		content: "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.commands)-1 {
				m.cursor++
			}
		case "enter", " ":
			command := m.commands[m.cursor]
			switch command {
			case "Paste from buffer":
				m.klipp.Write("newnote")
			case "List notes":
				m.content = strings.Join(m.klipp.GetNoteNames(), "\n")
			case "Read note":
				m.content = m.klipp.Read("test")
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	view := ""

	if m.content != "" {
		view += m.content + "\n"
	} else {
		view += "Commands\n\n"
		for i, command := range m.commands {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			view += fmt.Sprintf("%s %s\n", cursor, command)
		}
	}

	quitMessage := "\nPress q to quit.\n"
	view += quitMessage

	return view
}

func Start() {
	program := tea.NewProgram(initialModel())
	if err := program.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
