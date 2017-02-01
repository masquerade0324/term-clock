package main

import (
	"fmt"
	"time"

	termbox "github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	pollEvent()
}

func pollEvent() {
	draw()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			default:
				draw()
			}
		default:
			draw()
		}
	}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	fgColor := termbox.ColorRed
	bgColor := termbox.ColorWhite
	drawLine(10, 10, fmt.Sprintf("%v", time.Now()), fgColor, bgColor)

	termbox.Flush()
}

func drawLine(x, y int, str string, fg, bg termbox.Attribute) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		termbox.SetCell(x+i, y, runes[i], fg, bg)
	}
}
