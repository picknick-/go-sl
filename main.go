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

func DrawChart() {
	screen.Clear(text, background)
	screen.Flush()

}

func main() {
	err := screen.Init()
	check(err)
static.
	defer termbox.Close()

	for loop := true; loop; {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlQ, termbox.KeyF10:
				loop = false
				DrawChart()
			}
		case termbox.EventResize:
			DrawChart()
		}
	}
}
