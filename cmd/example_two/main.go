package main

import (
    "fmt"
    "os"
    "time"
    "net/http"

    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    status int
    message string
    err error
}

func (m model) Init() tea.Cmd {
    return checkServer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String(){
            case "q":
                return m, tea.Quit
        }
    case errMsg:
        m.err = msg
        return m, nil
    case statusMsg:
        m.message = "Omfg it worked"
        m.status = int(msg)
        return m, nil
    }
    return m, nil
}

func (m model) View() string {
    return fmt.Sprintf("%s%d", m.message, m.status)
}

func checkServer() tea.Msg {
    c := &http.Client{Timeout: 10 * time.Second}
    res, err := c.Get("https://charm.sh/")
    time.Sleep(5*time.Second)
    if err != nil {
        return errMsg{err}
    }
    return statusMsg(res.StatusCode)
}

type statusMsg int
type errMsg struct { err error  }

func (e errMsg) Error() string { return e.err.Error() }

func main() {
    m := model{message: "Making request"}
    _, err := tea.NewProgram(m).Run()
    if err != nil {
        fmt.Printf("Error creating new program: %v", err)
        os.Exit(-1)
    }
}
