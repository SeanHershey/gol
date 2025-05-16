package core

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

func HelpBox() string {
	dialogBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#71797E")).
		Padding(1, 3, 1, 2).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
	options := list.New(
		"h Show help",
		"q Quit",
		"b Blank canvas",
		"r Random canvas",
		"j Prev algo",
		"k Next algo",
		"c Cycle colors",
		"[space] Pause",
		"[click] toggle cell",
	).Enumerator(list.Dash)
	return dialogBoxStyle.Render(options.String())
}
