package main

import (
	"time"

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
	for i, section := range train {
		printString(0, i, section)
	}
	printString(0, 8, "KAKAOWY POCIĄG")
	screen.Flush()

}

func drawTrain(x, y int) {
	screen.Clear(text, background)
	wheel := x % 5
	for i, section := range append(train, getWheels(wheel)...) {
		printString(x, i+y, section)
	}

	screen.Flush()
}

func getWheels(a int) []string {
	switch a {
	case 0:
		return wheels0
	case 1:
		return wheels1
	case 2:
		return wheels2
	case 3:
		return wheels3
	case 4:
		return wheels4
	case 5:
		return wheels5
	}
	return wheels0
}

func printString(x, y int, s string) {
	offsetX := 0
	offsetY := 0
	for _, char := range s {
		if x+offsetX > screen.Width-1 {
			return
		}
		if x+offsetX > 0 {
			screen.Buffer[x+offsetX+(y+offsetY)*screen.Width].Ch = char
		}
		offsetX++
	}
}

func main() {
	err := screen.Init()
	check(err)
	defer termbox.Close()
	ticker := time.NewTicker(40000 * time.Microsecond)
	done := make(chan bool)
	trainOffset := screen.Width

	go func() {
		for {
			<-ticker.C
			trainOffset--
			drawTrain(trainOffset, screen.Height/2-5)
			if trainOffset == -len(train[0]) {
				done <- true
			}
		}
	}()

	<-done
}
