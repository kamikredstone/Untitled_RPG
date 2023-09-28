package main

import (
	"errors"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/kamikredstone/Untitled_RPG/entities"
	"github.com/kamikredstone/Untitled_RPG/rendering"
	"github.com/kamikredstone/Untitled_RPG/world"
	"github.com/rivo/tview"
)

func moveEntity(e entities.Entity, x int, y int, room *world.Room) error {
	if x < 0 || x >= len(room.Map.Tiles[0]) || y < 0 || y >= len(room.Map.Tiles) {
		return errors.New("out of bounds")
	}
	if !(*room).Map.Tiles[y][x].TerrainType.IsWalkable {
		return errors.New("can't move to the specified tile")
	}
	// check if the destination is a door
	doorKey := fmt.Sprintf("%d,%d", x, y)
	if door, exists := room.Doors[doorKey]; exists {
		room = door.ToRoom
		x, y = door.DestinationEntry["X"], door.DestinationEntry["Y"]
	}
	oldX, oldY := e.GetPosition()["X"], e.GetPosition()["Y"]
	room.Map.Tiles[oldY][oldX].Entity = nil // remove the player from the old position
	err := e.SetPosition(x, y)
	if err == nil {
		room.Map.Tiles[y][x].Entity = e // set the player in the new position
		return nil
	}
	return err
}

func main() {

	// Render map
	// initialize renderer
	renderer := rendering.AsciiRenderer{}
	// Create room
	grassTerrain := "·"
	terrain := world.GetTerrain(1, "grass", grassTerrain, grassTerrain, true)
	worldMap := world.CreateMapSingleTerrain(terrain, 40, 20)
	border := world.CreateBorder("|", "_", "‾")
	borderedMap := world.AddBorder(*worldMap, border)
	mainRoom := world.CreateRoom(borderedMap)
	fmt.Print(renderer.RenderRoom(&mainRoom))

	app := tview.NewApplication()
	newPlayerPosition := make(map[string]int)
	newPlayerPosition["X"] = 1
	newPlayerPosition["Y"] = 1
	player := entities.NewPlayer("Gil", "@", newPlayerPosition)
	mainRoom.Map.Tiles[newPlayerPosition["X"]][newPlayerPosition["Y"]].Entity = player

	textView := tview.NewTextView().SetScrollable(false).SetText(renderer.RenderRoom(&mainRoom))

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			moveEntity(player, player.Position["X"], player.Position["Y"]-1, &mainRoom)
		case tcell.KeyDown:
			moveEntity(player, player.Position["X"], player.Position["Y"]+1, &mainRoom)
		case tcell.KeyLeft:
			moveEntity(player, player.Position["X"]-1, player.Position["Y"], &mainRoom)
		case tcell.KeyRight:
			moveEntity(player, player.Position["X"]+1, player.Position["Y"], &mainRoom)
		}
		textView.SetText(renderer.RenderRoom(&mainRoom))
		return event
	})

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
