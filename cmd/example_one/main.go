package main

import (
    "fmt"
    "os"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    message string
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type){
    case tea.KeyMsg:
        switch msg.String(){
        case "q":
            return m, tea.Quit
        }
    }
    
    return m, nil
}

func (m model) View() string {
    return m.message
}

func main() {
    m := model{
        message: "Waiting for death",
    }
    _, err := tea.NewProgram(m).Run()
    if err != nil {
        fmt.Printf("Error starting %v", err)
        os.Exit(-1)
    }
}
