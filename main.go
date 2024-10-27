package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
	"strings"
)

const github = "https://github.com"
const github_rr = "https://github.com/remoterabbit"

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Space key.Binding
	Enter key.Binding
	Help  key.Binding
	Quit  key.Binding
}

type repository struct {
	name        string
	url         string
	description string
	selected    bool
}

type model struct {
	keys     keyMap
	help     help.Model
	quitting bool
	repos    []repository
	cursor   int
	selected map[int]bool
}

var repos = []repository{
	{"Dot_Files", fmt.Sprintf("%s/dot_files", github_rr), "All the dot files (kind of)", false},
	{"Tmux", fmt.Sprintf("%s/tmux", github_rr), "Dots for tmux setup.", false},
	{"zsh", fmt.Sprintf("%s/zsh", github_rr), "Dots for zsh setup.", false},
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Space, k.Enter},
		{k.Help, k.Quit},
	}
}

var DefaultKeyMap = keyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "move along"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Space: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "select"),
	),
}

func initialModel() model {
	return model{
		help:     help.New(),
		keys:     DefaultKeyMap,
		repos:    repos,
		selected: make(map[int]bool),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.keys.Down):
			if m.cursor < len(m.repos)-1 {
				m.cursor++
			}
		case key.Matches(msg, m.keys.Space):
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = true
			}
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Enter):
			dirname, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}

			for i, repo := range m.repos {
				dir := fmt.Sprintf("%s/repos/personal/clone/%s", dirname, repo.name)
				if _, selected := m.selected[i]; selected {
					fmt.Printf("Selected repository: %s at %s\n", repo.name, repo.url)
					_, err := git.PlainClone(dir, false, &git.CloneOptions{
						URL: repo.url,
					})
					if err != nil {
						fmt.Printf("Error running PLL: %v", err)
					}
				}
			}
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF69B4")).
			Bold(true).
			Align(lipgloss.Center).
			MarginBottom(1)
	listStyle = lipgloss.NewStyle().
			MarginLeft(2)
)

func (m model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	status := titleStyle.Render("PC Load Letter") + "\n"

	for i, repo := range m.repos {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := "[ ]"
		if _, ok := m.selected[i]; ok {
			checked = "[X]"
		}
		line := fmt.Sprintf("%s %s %s - %s", cursor, checked, repo.name, repo.description)
		status += listStyle.Render(line) + "\n"
	}

	helpView := m.help.View(m.keys)
	height := 8 - strings.Count(status, "\n") - strings.Count(helpView, "\n")

	return "\n" + status + strings.Repeat("\n", height) + helpView
}

func main() {
	p := tea.NewProgram(initialModel())
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Error running PLL: %v", err)
	}

	// finalModel := m.(model)
}
