package main

// A simple example demonstrating how to draw and animate on a cellular grid.
// Note that the cellbuffer implementation in this example does not support
// double-width runes.

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	fps = 20
)

var (
	blue   color.Color = color.RGBA{69, 145, 196, 255}
	yellow color.Color = color.RGBA{255, 230, 120, 255}
	orange             = lipgloss.NewStyle()
)

type cellbuffer struct {
	grid   [][]uint8
	buffer [][]uint8
	width  int
	height int
	cx     map[int]bool
	cy     map[int]bool
}

func (c *cellbuffer) update() {
	for _y := range c.height - 2 {
		for _x := range c.width - 2 {
			if c.cx[_x] && c.cy[_y] {
				c.buffer[_x][_y+1] = 1
				continue
			}
			y := _y + 1
			x := _x + 1
			c.buffer[x][y] = 0

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
	}

	c.cx = map[int]bool{}
	c.cy = map[int]bool{}
	temp := c.buffer
	c.buffer = c.grid
	c.grid = temp
}

func (c *cellbuffer) init(w, h int) {
	if w == 0 {
		return
	}
	c.grid = make([][]uint8, w)
	c.buffer = make([][]uint8, w)
	for i := range c.grid {
		c.grid[i] = make([]uint8, h)
		c.buffer[i] = make([]uint8, h)
	}
	c.width = w
	c.height = h
	c.reset()
}

func (c *cellbuffer) reset() {
	for _y := range c.height - 2 {
		for _x := range c.width - 2 {
			x := _x + 1
			y := _y + 1
			if rand.Float32() < 0.5 {
				c.grid[x][y] = 1
			} else {
				c.grid[x][y] = 0
			}
		}
	}
}

func (c cellbuffer) ready() bool {
	return len(c.buffer) > 0
}

func (c cellbuffer) String() string {
	var b strings.Builder
	for y := range c.height {
		for x := range c.width {
			switch c.grid[x][y] {
			case 0:
				b.WriteRune(' ')
			case 1:
				c, _ := colorful.Hex("#FF5733")
				s := lipgloss.NewStyle().SetString(" ").Background(lipgloss.Color(c.Hex()))
				b.WriteString(s.String())
			case 2:
				c, _ := colorful.Hex("#FFFFFF")
				s := lipgloss.NewStyle().SetString(" ").Background(lipgloss.Color(c.Hex()))
				b.WriteString(s.String())
			default:
				b.WriteRune(' ')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

type frameMsg struct{}

func animate() tea.Cmd {
	return tea.Tick(time.Second/fps, func(_ time.Time) tea.Msg {
		return frameMsg{}
	})
}

type model struct {
	cells cellbuffer
}

func (m model) Init() tea.Cmd {
	return animate()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tea.WindowSizeMsg:
		m.cells.init(msg.Width, msg.Height)
		return m, nil
	case tea.MouseMsg:
		m.cells.cx[msg.X] = true
		m.cells.cy[msg.Y] = true
		return m, nil

	case frameMsg:
		m.cells.update()
		return m, animate()
	default:
		return m, nil
	}
}

func (m model) View() string {
	return m.cells.String()
}

func main() {
	m := model{}

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
}
