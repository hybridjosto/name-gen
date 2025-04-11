package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hybridjosto/namegen-go/lib"
)

type model struct {
	name     string
	gender   string
	quitting bool
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.quitting = true
			return m, tea.Quit
		case "m":
			m.gender = "male"
			m.name = lib.GenerateName("male")
		case "f":
			m.gender = "female"
			m.name = lib.GenerateName("female")
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true).
		Align(lipgloss.Center).
		Render("✨ Name Generator Supreme ✨")

	nameStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("42")).
		Bold(true).
		Italic(true).
		Underline(true).
		Align(lipgloss.Center)

	nameBox := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Margin(1, 0).
		Align(lipgloss.Center).
		Render(nameStyle.Render(m.name))

	footer := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true).
		Align(lipgloss.Center).
		Render("[m] Male   [f] Female   [q] Quit")

	return fmt.Sprintf("%s\n%s\n%s", title, nameBox, footer)
}

func RunMaxMode() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running interactive mode:", err)
		os.Exit(1)
	}
}
