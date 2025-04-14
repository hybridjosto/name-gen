package ui

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hybridjosto/namegen-go/lib"
)

type model struct {
	name       string
	gender     string
	quitting   bool
	addedNames []string
	scroll     int
	statusMsg  string
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

var addedNames []string

var tableStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(1, 2).
	Width(30).
	Height(10) // change based on your layout

func renderTable(names []string, scroll int, height int) string {
	end := scroll + height
	if end > len(names) {
		end = len(names)
	}

	visible := names[scroll:end]

	rows := make([]string, len(visible))
	for i, name := range visible {
		rows[i] = fmt.Sprintf("%2d. %s", scroll+i+1, name)
	}

	return tableStyle.Render(strings.Join(rows, "\n"))
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
		case "a":
			if m.name == "" {
				m.statusMsg = "⚠️ No name to add"
			} else if !contains(m.addedNames, m.name) {
				m.addedNames = append(m.addedNames, m.name)
				lib.WriteToFile(m.name)
				m.statusMsg = "✅ Added!"

			} else {
				m.statusMsg = "⚠️ Already in list"
			}
		case "j", "down":
			if m.scroll < len(m.addedNames)-1 {
				m.scroll++
			}
		case "k", "up":
			if m.scroll > 0 {
				m.scroll--
			}
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
		Render("[m] Male   [f] Female   [a] Add   [j] Scroll Down   [k] Scroll Up   [q] Quit")

	// Left side: title, name box, footer
	left := lipgloss.JoinVertical(lipgloss.Top, title, nameBox, footer, m.statusMsg)

	// Right side: scrollable name list
	right := renderTable(m.addedNames, m.scroll, 10) // 10 = visible rows

	// Final layout: left + right side-by-side
	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

func RunMaxMode() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running interactive mode:", err)
		os.Exit(1)
	}
}
