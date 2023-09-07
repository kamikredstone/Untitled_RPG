package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/kamikredstone/Untitled_RPG/rendering"
	"github.com/kamikredstone/Untitled_RPG/world"
	"github.com/rivo/tview"
)

func main() {

	// Render map
	// initialize renderer
	renderer := rendering.AsciiRenderer{}
	// Create room
	terrain := world.GetTerrain(1, "grass", graphic, isWalkable)

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
