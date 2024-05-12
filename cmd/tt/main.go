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
var correctStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#04b513")).Background(lipgloss.Color("#a8f0ae"))
var inCorrectStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f72f43")).Background(lipgloss.Color("#fa8c97"))
var cursorStyle = lipgloss.NewStyle().Background(lipgloss.Color("#555555"))

func initModel() model {
    sentence := sentences[rand.Intn(len(sentences))]
    userInput := make([]lipgloss.Style, len(sentence))
    userInput[0] = cursorStyle
    for i := range userInput[1:] {
        userInput[i+1] = remainingSentenceStyle
    }
    return model{
        sentence: sentence,
        userInput: userInput,
    }
}

type model struct {
    sentence string
    userInput []lipgloss.Style
    cursor int
}

func (m *model) progressCursor(letter byte) {
    if letter == m.sentence[m.cursor]{
        m.userInput[m.cursor] = correctStyle
    } else {
        m.userInput[m.cursor] = inCorrectStyle
    }
    if m.cursor < len(m.sentence)-1 {
        m.cursor++
        m.userInput[m.cursor] = cursorStyle 
    }
}

func (m *model) decrementCursor() {
    if m.cursor > 0 {
        m.userInput[m.cursor] = remainingSentenceStyle
        m.cursor--
        m.userInput[m.cursor] = cursorStyle
    }
}
    

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type){
    case tea.KeyMsg:
        switch msg.Type{
            case tea.KeyRunes:
                if len(msg.Runes) == 1 {
                    m.progressCursor(byte(msg.Runes[0]))
                }
            case tea.KeySpace:
                m.progressCursor(byte(' '))
            case tea.KeyBackspace:
                m.decrementCursor()
           case tea.KeyCtrlC:
                return m, tea.Quit
        }
    }


    return m, nil // tea.Batch(taCmd)
}

func (m model) View() string {
    var s strings.Builder
    for i, renderer := range m.userInput {
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
