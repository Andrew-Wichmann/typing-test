package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

    "github.com/Andrew-Wichmann/typing-test/pkg/textTest"
)

var welcomeMessage = "Hello! Welcome to the typing test challenge! When you're ready, press <enter> to begin"
var finishedMessage = "Test finished! Press <enter> to go again"   

type model struct {
    test textTest.Model
    startPage bool
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    if val, ok := msg.(tea.KeyMsg); ok {
        if val.Type == tea.KeyCtrlC {
            return m, tea.Quit
        }
    }
    if m.startPage {
        if val, ok := msg.(tea.KeyMsg); ok {
            if val.Type == tea.KeyEnter {
                m.startPage = false
            }
        }
        return m, nil
    }
    if m.test.Done {
        if val, ok := msg.(tea.KeyMsg); ok {
            if val.Type == tea.KeyEnter {
                m.test = textTest.NewModel()
            }
        }
        return m, nil
    }   

    test, cmd := m.test.Update(msg)
    m.test = test

    return m, cmd
}

func (m model) View() string {
    if m.startPage {
        return welcomeMessage
    }
    if m.test.Done {
        return finishedMessage
    }
    return m.test.View()
} 

func main() {
    m := model{
        test: textTest.NewModel(),
        startPage: true,
    }
    _, err := tea.NewProgram(m).Run()
    if err != nil {
        fmt.Printf("Error creating new program: %v", err)
        os.Exit(-1)
    }
}
