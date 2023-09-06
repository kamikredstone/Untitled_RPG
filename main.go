package main

import (
	"github.com/kamikredstone/Untitled_RPG/entities"
	"github.com/kamikredstone/Untitled_RPG/rendering"
	"github.com/kamikredstone/Untitled_RPG/world"
)

func main() {
	// Generate Map
	terrain := world.GetTerrain(1, "grass", "·", true)
	worldMap := world.CreateMapSingleTerrain(terrain, 40, 20)
	border := world.CreateBorder("|", "_", "‾")
	borderedMap := world.AddBorder(*worldMap, border)
	mainRoom := world.CreateRoom(borderedMap)
	// Render room
	renderer := rendering.AsciiRenderer{}
	character := entities.CreateDummyPlayerCharacter(&mainRoom)
	renderer.RenderRoom(mainRoom)
}
