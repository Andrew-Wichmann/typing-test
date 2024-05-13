package textTest

import (
    "math/rand"
    "strings"

    tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


var sentences = []string{
    "Mary had a little lamb. Little lamb. Little lamb.",
    "The quick brown fox jumped over the lazy dog.",
    "This is an example sentence",
}

var windowStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Align(lipgloss.Center, lipgloss.Center)

var baseStyle = lipgloss.NewStyle()
var remainingSentenceStyle = lipgloss.NewStyle().Inherit(baseStyle).Faint(true)
var correctStyle = lipgloss.NewStyle().Inherit(baseStyle).Foreground(lipgloss.Color("#04b513")).Background(lipgloss.Color("#a8f0ae"))
var inCorrectStyle = lipgloss.NewStyle().Inherit(baseStyle).Foreground(lipgloss.Color("#f72f43")).Background(lipgloss.Color("#fa8c97")).Strikethrough(true)
var cursorStyle = lipgloss.NewStyle().Inherit(baseStyle).Background(lipgloss.Color("#555555")).Blink(true)

type Model struct {
    sentence string
    userInput []lipgloss.Style
    cursor int
    Done bool
}

func (m *Model) progressCursor(letter byte) {
    if letter == m.sentence[m.cursor]{
        m.userInput[m.cursor] = correctStyle
    } else {
        m.userInput[m.cursor] = inCorrectStyle
    }
    if m.cursor < len(m.sentence)-1 {
        m.cursor++
        m.userInput[m.cursor] = cursorStyle 
    } else {
        m.Done = true
    }
}

func (m *Model) decrementCursor() {
    if m.cursor > 0 {
        m.userInput[m.cursor] = remainingSentenceStyle
        m.cursor--
        m.userInput[m.cursor] = cursorStyle
    }
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
    if m.Done {
        return m, nil
    }
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
       }
    case tea.WindowSizeMsg:
        windowStyle.Width(msg.Width).Height(msg.Height-2)
    }
    return m, nil
}

func (m Model) View() string {
    var s strings.Builder
    for i, renderer := range m.userInput {
        s.WriteString(renderer.Render(string(m.sentence[i])))
    }
    return windowStyle.Render(s.String())
}

func NewModel() Model {
    m := Model{}
    sentence := sentences[rand.Intn(len(sentences))]
    userInput := make([]lipgloss.Style, len(sentence))
    userInput[0] = cursorStyle
    for i := range userInput[1:] {
        userInput[i+1] = remainingSentenceStyle
    }
    m.sentence = sentence
    m.userInput = userInput
    m.cursor = 0
    return m 
}

