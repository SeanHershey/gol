package main

// A simple example demonstrating how to draw and animate on a cellular grid.
// Note that the cellbuffer implementation in this example does not support
// double-width runes.

import (
	"fmt"
	"image/color"
	"os"

	"gol/core"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	fps = 20
)

var (
	blue   color.Color = color.RGBA{69, 145, 196, 255}
	yellow color.Color = color.RGBA{255, 230, 120, 255}
	orange             = lipgloss.NewStyle()
)

type model struct {
	Core core.Core
}

func (m model) Init() tea.Cmd {
	return m.Core.Animate()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, m.Core.HandleKeyMsg(msg)
	case tea.WindowSizeMsg:
		m.Core.Init(msg.Width, msg.Height)
		m.Core.Random()
		return m, nil
	case tea.MouseMsg:
		return m, m.Core.HandleMouseMsg(msg)
	case core.FrameMsg:
		m.Core.Update()
		return m, m.Core.Animate()
	default:
		return m, nil
	}
}

func (m model) View() string {
	return m.Core.String()
}

func main() {
	m := model{}
	m.Core.Fps = fps

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
}
