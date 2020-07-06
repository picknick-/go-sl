package main

import (
	"github.com/nsf/termbox-go"
)

//Screen Type for handling screen
type Screen struct {
	Buffer []termbox.Cell
	Width  int
	Height int
}

//reScreen
func (s *Screen) reScreen() {
	s.Buffer = termbox.CellBuffer()
	s.Width, s.Height = termbox.Size()
}

//Init Screen
func (s *Screen) Init() error {
	err := termbox.Init()
	if err == nil {
		s.reScreen()
	}
	return err
}

//Flush screen
func (s *Screen) Flush() {
	termbox.Flush()
	s.reScreen()
}

//Clear screen
func (s *Screen) Clear(fg, bg termbox.Attribute) {
	termbox.Clear(fg, bg)
	s.reScreen()
}

var (
	text, background termbox.Attribute
	lightness        int
	helpVisible      bool
	screen           Screen
)

func termNative(c int) termbox.Attribute {
	return termbox.Attribute(c + 1)
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func fill(x1, y1, x2, y2 int, fg, bg termbox.Attribute) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			termbox.SetCell(x, y, ' ', fg, bg)
		}
	}
}

func drawChart() {
	screen.Clear(text, background)
	fill(0, 0, screen.Width-1, screen.Height-1, background, termbox.ColorRed)
	printString(0, 2, d51STR1)
	printString(0, 3, d51STR2)
	printString(0, 8, "KAKAOWY POCIĄg")
	screen.Flush()

}

func printString(x, y int, s string) {
	offsetX := 0
	offsetY := 0
	for _, char := range s {
		if char == '\n' {
			offsetX = 0
			offsetY++
		} else {
			screen.Buffer[x+offsetX+(y+offsetY)*screen.Width].Ch = char
			screen.Buffer[x+offsetX+(y+offsetY)*screen.Width].Fg = termbox.ColorBlue
			offsetX++
		}
	}
}

func main() {
	err := screen.Init()
	check(err)

	defer termbox.Close()
	drawChart()
	for loop := true; loop; {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlQ, termbox.KeyF10:
				loop = false
				drawChart()
			}
		case termbox.EventResize:
			drawChart()
		}
	}
}
