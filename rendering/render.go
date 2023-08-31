package rendering

import (
	"fmt"

	"github.com/kamikredstone/Untitled_RPG/world"
)

type AsciiRenderer struct{} // Implements the AsciiRenderer interface

func (a *AsciiRenderer) RenderTile(t world.Tile) string {
	return t.TerrainType.Graphic
}

func (a *AsciiRenderer) RenderMap(m world.Map) {
	for _, row := range m.Tiles {
		for _, tile := range row {
			fmt.Print(a.RenderTile(tile))
		}
		fmt.Println()
	}
}
