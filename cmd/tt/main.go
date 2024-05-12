package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

    "github.com/Andrew-Wichmann/typing-test/pkg/textTest"
)

type model struct {
    test textTest.Model
}


func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m.test.Update(msg)
}

func (m model) View() string {
    return m.test.View()
}

func main() {
    m := model{
        test: textTest.NewModel(),
    }
    _, err := tea.NewProgram(m).Run()
    if err != nil {
        fmt.Printf("Error creating new program: %v", err)
        os.Exit(-1)
    }
}
