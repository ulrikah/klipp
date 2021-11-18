package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"golang.design/x/clipboard"
)

// START CLI

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func initialModel() model {
	return model{
		// Our shopping list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

// END CLI

type Klipp struct {
	HomeDir string
}

func (k Klipp) pathTo(relPath string) string {
	// TODO: probably some util in ioutils to concat paths
	return k.HomeDir + "/" + relPath
}

func (k Klipp) Read(noteName string) string {
	// Reads from the specified note and copies it to the clipboard

	content, err := ioutil.ReadFile(k.pathTo(noteName))
	if err != nil {
		return ""
	} else {
		copyToClipboard(string(content))
		return string(content)
	}
}

func (k Klipp) Write(note string) string {
	// Writes to a specified note from the clipboard

	// TODO: properly handle errors
	absPath := k.pathTo(note)
	_, err := ioutil.ReadFile(absPath)
	if err != nil {
		ioutil.WriteFile(absPath, []byte(pasteFromClipboard()), 0644)
		pasteFromClipboard()
		return "success"
	}
	return "failure"
}

func makeTempFile(name string) *os.File {
	tmpfile, err := ioutil.TempFile("", name)
	if err != nil {
		log.Fatal(err)
	}
	return tmpfile
}

func pasteFromClipboard() string {
	return string(clipboard.Read(clipboard.FmtText))
}

func copyToClipboard(msg string) <-chan struct{} {
	return clipboard.Write(clipboard.FmtText, []byte(msg))
}

func main() {
	// k := Klipp{HomeDir: "/Users/ulrikah/.klipp"}
	// k.Write("test")

	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	// log.Printf("In the clipboard: %s", pasteFromClipboard())
	// for i := 0; i < 10; i++ {
	// 	copyToClipboard(fmt.Sprintf("%s%d", "Hello: ", i))
	// 	log.Printf("In the clipboard: %s", pasteFromClipboard())
	// }
	// files, err := ioutil.ReadDir("./")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, file := range files {
	// 	log.Println(file.Name(), "is directory?", file.IsDir())
	// }
	// tmpfile := makeTempFile("hush")
	// defer os.Remove(tmpfile.Name()) // clean up

	// message := []byte("\n\n\tHello World\n\n")
	// if _, err := tmpfile.Write(message); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := tmpfile.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	// content, err := ioutil.ReadFile(tmpfile.Name())
	// log.Printf("File contents: %s", content)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
