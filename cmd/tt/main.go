package main

import (
    "fmt"
    "os"

    "github.com/charmbracelet/bubbles/textarea"
    "github.com/charmbracelet/bubbles/viewport"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

func initModel() model {
    ta := textarea.New()
    ta.Placeholder = "Foobar"
    ta.Prompt = "┃ "
    ta.Focus()
    
    // Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

    return model{textarea: ta}
}

type model struct {
    viewport viewport.Model
    messages []string
    textarea textarea.Model
}

func (m model) Init() tea.Cmd {
    return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var taCmd tea.Cmd
    m.textarea, taCmd = m.textarea.Update(msg)

    switch msg := msg.(type){
    case tea.KeyMsg:
        switch msg.String(){
        case "q":
            return m, tea.Quit
        }
    }
    return m, tea.Batch(taCmd)
}

func (m model) View() string {
    return m.textarea.View()
}

func main() {
    m := initModel()
    _, err := tea.NewProgram(m).Run()
    if err != nil {
        fmt.Printf("Error creating new program: %v", err)
        os.Exit(-1)
    }
}
