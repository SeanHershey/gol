package core

import tea "github.com/charmbracelet/bubbletea"

func (c *Core) HandleMouseMsg(msg tea.MouseMsg) tea.Cmd {
	switch msg.Action {
	case tea.MouseActionPress, tea.MouseActionMotion:
		if c.grid[msg.X][msg.Y] == 1 {
			c.grid[msg.X][msg.Y] = 0
			return nil
		}
		c.grid[msg.X][msg.Y] = 1
		return nil
	default:
		return nil
	}
}
