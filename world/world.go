package world

import (
	"github.com/google/uuid"
)

type Renderer interface {
	RenderTile(t Tile) string
	RenderMap(m Map)
}

type Terrain struct {
	MovementMultiplier int
	Name               string
	Graphic            string
	IsWalkable         bool
}

type Tile struct {
	TerrainType Terrain
}

type Map struct {
	Tiles  [][]Tile
	Size_X int
	Size_Y int
}

type Room struct {
	Map  Map
	GUID uuid.UUID
}

type Border struct {
	VerticalGraphic         string
	HorizontalTopGraphic    string
	HorizontalBottomGraphic string
}

func CreateMapSingleTerrain(terrainType Terrain, size_x int, size_y int) *Map {
	newMap := &Map{
		Tiles:  make([][]Tile, size_x),
		Size_X: size_x,
		Size_Y: size_y,
	}
	for i := range newMap.Tiles {
		newMap.Tiles[i] = make([]Tile, size_y)
		for j := range newMap.Tiles[i] {
			newMap.Tiles[i][j] = GetTile(terrainType)
		}
	}
	return newMap
}

func CreateBorder(verticalGraphic string, horizontalTopGraphic string, horizontalBottomGraphic string) Border {
	return Border{
		VerticalGraphic:         verticalGraphic,
		HorizontalTopGraphic:    horizontalTopGraphic,
		HorizontalBottomGraphic: horizontalBottomGraphic,
	}
}

func GetTerrain(mvmnt int, name string, graphic string, isWalkable bool) Terrain {
	return Terrain{
		MovementMultiplier: mvmnt,
		Name:               name,
		Graphic:            graphic,
		IsWalkable:         isWalkable,
	}
}

func GetTile(terrainType Terrain) Tile {
	return Tile{TerrainType: terrainType}
}

func AddBorder(originalMap Map, border Border) Map {
	if len(originalMap.Tiles) == 0 {
		return originalMap
	}

	xLen := originalMap.Size_X + 2
	yLen := originalMap.Size_Y + 2
	originalTile := originalMap.Tiles[0][0]
	newMap := Map{
		Tiles:  make([][]Tile, yLen),
		Size_X: xLen,
		Size_Y: yLen,
	}

	// Create top border tile
	topTerrain := GetTerrain(0, "border", border.HorizontalTopGraphic, false)
	topTile := GetTile(topTerrain)
	// Create bottom border tile
	bottomTerrain := GetTerrain(0, "border", border.HorizontalBottomGraphic, false)
	bottomTile := GetTile(bottomTerrain)
	// Create vertical border tile
	verticalTerrain := GetTerrain(0, "border", border.VerticalGraphic, false)
	verticalTile := GetTile(verticalTerrain)

	// Fill the new map
	for row := 0; row <= yLen-1; row++ {
		newMap.Tiles[row] = make([]Tile, xLen)
		if row == 0 {
			//fill with top border tiles
			for i := 0; i < xLen; i++ {
				newMap.Tiles[row][i] = topTile
			}
		} else if row == yLen-1 {
			//fill with bottom border tiles
			for i := 0; i < xLen; i++ {
				newMap.Tiles[row][i] = bottomTile
			}
		} else {
			for i := 0; i < xLen; i++ {
				if i == 0 || i == xLen-1 {
					newMap.Tiles[row][i] = verticalTile
				} else {
					newMap.Tiles[row][i] = originalTile
				}
			}
		}
	}
	return newMap
}
