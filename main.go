package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	playerChar = '@'
	emptySpace = ' '
	mapWidth   = 20
	mapHeight  = 20
)

var (
	playerX = mapWidth / 2
	playerY = mapHeight / 2
	gameMap [mapHeight][mapWidth]rune
)

func main() {
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			gameMap[y][x] = emptySpace
		}
	}

	app := tview.NewApplication()

	textView := tview.NewTextView().SetText(drawMap())

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			if playerY > 0 {
				playerY--
			}
		case tcell.KeyDown:
			if playerY < mapHeight-1 {
				playerY++
			}
		case tcell.KeyLeft:
			if playerX > 0 {
				playerX--
			}
		case tcell.KeyRight:
			if playerX < mapWidth-1 {
				playerX++
			}
		}
		textView.SetText(drawMap())
		return event
	})

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
