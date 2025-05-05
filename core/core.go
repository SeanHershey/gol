package core

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/mattn/go-runewidth"
)

type Core struct {
	grid     [][]uint8
	buffer   [][]uint8
	width    int
	height   int
	paused   bool
	Fps      int
	hideHelp bool
}

type FrameMsg struct{}

func (c *Core) Animate() tea.Cmd {
	if c.paused {
		return nil
	}
	return tea.Tick(time.Second/time.Duration(c.Fps), func(_ time.Time) tea.Msg {
		return FrameMsg{}
	})
}

func (c *Core) Reset() {
	w := c.width
	h := c.height
	c.grid = make([][]uint8, w)
	c.buffer = make([][]uint8, w)
	for i := range c.grid {
		c.grid[i] = make([]uint8, h)
		c.buffer[i] = make([]uint8, h)
	}
}

func (c *Core) Init(w, h int) {
	if w == 0 {
		return
	}
	c.width = w
	c.height = h
	c.Reset()
}

func (c *Core) Update() {
	if !c.Ready() {
		return
	}
	for _y := range c.height - 2 {
		for _x := range c.width - 2 {
			y := _y + 1
			x := _x + 1
			c.buffer[x][y] = 0
			c.SingleCell(x, y)
			// c.TwoCell(x, y)
		}
	}

	temp := c.buffer
	c.buffer = c.grid
	c.grid = temp
}

func (c *Core) Random() {
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

func (c Core) Ready() bool {
	return len(c.grid) > 0
}

func removeColorFromString(text string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(text, "")
}

func (c Core) String() string {
	var help string
	if !c.hideHelp {
		help = HelpBox()
	}
	hlines := strings.Split(help, "\n")
	var b strings.Builder
	startBoxX := c.width - runewidth.StringWidth(removeColorFromString(hlines[0]))
	endBoxY := len(hlines)
	for y := range c.height {
		for x := range c.width {
			if y < endBoxY && x >= startBoxX {
				b.WriteString(hlines[y])
				continue
			}
			switch c.grid[x][y] {
			case 0:
				b.WriteRune(' ')
			case 1:
				c, _ := colorful.Hex("#FFFFFF")
				s := lipgloss.NewStyle().SetString(" ").Background(lipgloss.Color(c.Hex()))
				b.WriteString(s.String())
			case 2:
				c, _ := colorful.Hex("#FF5733")
				s := lipgloss.NewStyle().SetString(" ").Background(lipgloss.Color(c.Hex()))
				b.WriteString(s.String())
			default:
				b.WriteRune(' ')
			}
		}
		if y == c.height-1 {
			continue
		}
		b.WriteRune('\n')
	}
	return b.String()
}
