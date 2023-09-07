package rendering

import (
	"fmt"

	"github.com/kamikredstone/Untitled_RPG/world"
)

type AsciiRenderer struct{} // Implements the Renderer interface from world

func (a *AsciiRenderer) RenderTile(t world.Tile) string {
	return t.TerrainType.Graphic
}

func (a *AsciiRenderer) RenderRoom(m world.Room) {
	for _, row := range m.Map.Tiles {
		for _, tile := range row {
			fmt.Print(a.RenderTile(tile))
		}
		fmt.Println()
	}
}
