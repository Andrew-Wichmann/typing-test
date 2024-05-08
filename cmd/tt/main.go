package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var sentences = []string{
    "Mary had a little lamb. Little lamb. Little lamb.",
    "The quick brown fox jumped over the lazy dog.",
    "This is an example sentence",
}

var remainingSentenceStyle = lipgloss.NewStyle().Faint(true)
var correctStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))
var inCorrectStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
var cursorStyle = lipgloss.NewStyle().Background(lipgloss.Color("#555555"))

func initModel() model {
    sentence := sentences[rand.Intn(len(sentences))]
    answers := make([]lipgloss.Style, len(sentence))
    for i := range answers {
        answers[i] = remainingSentenceStyle
    }
    return model{
        sentence: sentence,
        answers: answers,
    }
}

type model struct {
    sentence string
    answers []lipgloss.Style
    cursor int
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

    switch msg := msg.(type){
    case tea.KeyMsg:
        switch msg.String(){
        case "ctrl+c":
            return m, tea.Quit
        case "backspace":
            if m.cursor > 0 {
                m.cursor--
            }
        }

    }
    return m, nil // tea.Batch(taCmd)
}

func (m model) View() string {
    var s strings.Builder
    for i, renderer := range m.answers {
        s.WriteString(renderer.Render(string(m.sentence[i])))
    }
    return s.String()
}

func main() {
    m := initModel()
    _, err := tea.NewProgram(m).Run()
    if err != nil {
        fmt.Printf("Error creating new program: %v", err)
        os.Exit(-1)
    }
}
