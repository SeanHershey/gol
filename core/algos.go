package core

func (c *Core) SingleCell(x, y int) {
	// Number of surrounding cells
	n := c.grid[x-1][y-1] +
		c.grid[x-1][y+0] +
		c.grid[x-1][y+1] +
		c.grid[x+0][y-1] +
		c.grid[x+0][y+1] +
		c.grid[x+1][y-1] +
		c.grid[x+1][y+0] +
		c.grid[x+1][y+1]

	switch {
	case c.grid[x][y] == 0 && n == 3:
		c.buffer[x][y] = 1
	case n < 2 || n > 3:
		c.buffer[x][y] = 0
	default:
		c.buffer[x][y] = c.grid[x][y]
	}
}

func (c *Core) TwoCell(x, y int) {
	// Number of surrounding cells
	n := c.grid[x-1][y-1] +
		c.grid[x-1][y+0] +
		c.grid[x-1][y+1] +
		c.grid[x+0][y-1] +
		c.grid[x+0][y+1] +
		c.grid[x+1][y-1] +
		c.grid[x+1][y+0] +
		c.grid[x+1][y+1]

	switch {
	case c.grid[x][y] == 0 && n == 3:
		c.buffer[x][y] = 1
	case n < 2:
		c.buffer[x][y] = 0
	case c.grid[x][y] == 1 && n > 3:
		c.buffer[x][y] = 2
	case n > 3:
		c.buffer[x][y] = 0
	default:
		c.buffer[x][y] = c.grid[x][y]
	}
}

