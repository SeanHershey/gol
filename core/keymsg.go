package core

import tea "github.com/charmbracelet/bubbletea"

func (c *Core) HandleKeyMsg(msg tea.KeyMsg) tea.Cmd {
	switch msg.String() {
	case " ":
		if c.paused {
			c.paused = false
			return c.Animate()
		}
		c.paused = true
		return nil
	case "h":
		c.hideHelp = !c.hideHelp
		return nil
	case "b":
		c.Reset()
		return nil
	case "r":
		c.Random()
		return nil
	default:
		return tea.Quit
	}
}
