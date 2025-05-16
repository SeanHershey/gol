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
	case "j":
		if c.algoIndex == 0 {
			c.algoIndex = NUM_ALGOS - 1
			return nil
		}
		c.algoIndex = (c.algoIndex - 1) % NUM_ALGOS
		return nil
	case "k":
		c.algoIndex = (c.algoIndex + 1) % NUM_ALGOS
		return nil
	case "c":
		c.colorIndex = (c.colorIndex + 1) % 5
		return nil
	case "q":
		return tea.Quit
	default:
		return nil
	}
}
