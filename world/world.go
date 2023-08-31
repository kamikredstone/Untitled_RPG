package world

type Renderer interface {
	RenderTile(t Tile) string
	RenderMap(m Map) string
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

func CreateMapSingleTerrain(terrainType Terrain, size_x int, size_y int) *Map {
	m := &Map{
		Tiles:  make([][]Tile, size_x),
		Size_X: size_x,
		Size_Y: size_y,
	}
	for i := range m.Tiles {
		m.Tiles[i] = make([]Tile, size_y)
		for j := range m.Tiles[i] {
			m.Tiles[i][j] = GetTile(terrainType)
		}
	}
	return m
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
