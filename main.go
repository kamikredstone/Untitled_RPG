package main

import (
	"github.com/kamikredstone/Untitled_RPG/rendering"
	"github.com/kamikredstone/Untitled_RPG/world"
)

func main() {
	// Generate Map
	terrain := world.GetTerrain(1, "grass", ".", true)
	worldMap := world.CreateMapSingleTerrain(terrain, 10, 20)
	renderer := rendering.AsciiRenderer{}
	renderer.RenderMap(*worldMap)

}
