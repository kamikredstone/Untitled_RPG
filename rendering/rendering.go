package rendering

import (
	"github.com/kamikredstone/Untitled_RPG/world"
)

type AsciiRenderer struct{} // Implements the Renderer interface from world

func (a *AsciiRenderer) RenderTile(t *world.Tile) string {
	if t.Entity != nil {
		return t.Entity.GetGraphic()
	}
	return t.TerrainType.Graphic
}

func (a *AsciiRenderer) RenderRoom(m *world.Room) string {
	var output string = ""
	for _, row := range m.Map.Tiles {
		for _, tile := range row {
			output += a.RenderTile(&tile)
		}
		output += "\n"
	}
	return output
}
