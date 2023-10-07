package world

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kamikredstone/Untitled_RPG/entities"
)

type Renderer interface {
	RenderTile(t Tile) string
	RenderMap(m Map)
}

// The most basic type of struct.
// The tile struct contains a Terrain struct.
type Terrain struct {
	MovementMultiplier int
	Name               string
	Graphic            string
	CurrentGraphic     string
	IsWalkable         bool
}

// The second basic type of struct.
// The Map struct contains a two-dimensional Tile struct.
// The tile struct can also represent an entity standing on it.
type Tile struct {
	TerrainType Terrain
	Entity      entities.Entity
}

// A struct that represents a door leading to a different room.
type Door struct {
	Graphic          string
	ToRoom           *Room
	DestinationEntry map[string]int // X: 5, Y: 2
}

// The third basic type of struct.
// The Room struct continas a Map struct.
// This struct represents the map of the room
// and the size of it.
type Map struct {
	Tiles  [][]Tile
	Size_X int
	Size_Y int
}

// The Room struct is the upper-most struct.
// Each instance of this struct represents a room in the map.
// It has a map of the room and a hash table of doors that connect
// to other rooms/levels.
type Room struct {
	Map   Map
	GUID  uuid.UUID
	Doors map[string]*Door //position of the door as key
}

// A struct that represents the Border graphics of the room.
type Border struct {
	VerticalGraphic         string
	HorizontalTopGraphic    string
	HorizontalBottomGraphic string
}

// This function is used to create a pointer to a map with a single terrain in it.
func CreateMapSingleTerrain(terrainType Terrain, size_x int, size_y int) *Map {
	newMap := &Map{
		Tiles:  make([][]Tile, size_x),
		Size_X: size_x,
		Size_Y: size_y,
	}
	for i := range newMap.Tiles {
		newMap.Tiles[i] = make([]Tile, size_y)
		for j := range newMap.Tiles[i] {
			newMap.Tiles[i][j] = CreateTile(terrainType)
		}
	}
	return newMap
}

// This function creates a border object.
func CreateBorder(verticalGraphic string, horizontalTopGraphic string, horizontalBottomGraphic string) Border {
	return Border{
		VerticalGraphic:         verticalGraphic,
		HorizontalTopGraphic:    horizontalTopGraphic,
		HorizontalBottomGraphic: horizontalBottomGraphic,
	}
}

// This function creates a Terrain object.
func CreateTerrain(mvmnt int, name string, graphic string, currentGraphic string, isWalkable bool) Terrain {
	return Terrain{
		MovementMultiplier: mvmnt,
		Name:               name,
		Graphic:            graphic,
		CurrentGraphic:     currentGraphic,
		IsWalkable:         isWalkable,
	}
}

// This function creates a Tile object.
func CreateTile(terrainType Terrain) Tile {
	return Tile{TerrainType: terrainType}
}

// This function adds a border to a Map.
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
	topTerrain := CreateTerrain(0, "border", border.HorizontalTopGraphic, border.HorizontalTopGraphic, false)
	topTile := CreateTile(topTerrain)
	// Create bottom border tile
	bottomTerrain := CreateTerrain(0, "border", border.HorizontalBottomGraphic, border.HorizontalBottomGraphic, false)
	bottomTile := CreateTile(bottomTerrain)
	// Create vertical border tile
	verticalTerrain := CreateTerrain(0, "border", border.VerticalGraphic, border.VerticalGraphic, false)
	verticalTile := CreateTile(verticalTerrain)

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

// This function adds doors to the border.
func AddDoorToRoom(originalRoom *Room, doorGraphic string, positionX int, positionY int, toRoom *Room) *Room {
	doorTerrain := CreateTerrain(1, "door", doorGraphic, doorGraphic, true)
	doorTile := CreateTile(doorTerrain)

	originalRoom.Map.Tiles[positionY][positionX] = doorTile

	if originalRoom.Doors == nil {
		originalRoom.Doors = make(map[string]*Door)
	}
	positionKey := fmt.Sprintf("%d,%d", positionX, positionY)
	originalRoom.Doors[positionKey] = &Door{
		Graphic: doorGraphic,
		ToRoom:  toRoom,
	}
	return originalRoom
}

// This function creates a room from a map.
func CreateRoom(roomMap Map) Room {
	return Room{
		Map:  roomMap,
		GUID: uuid.New(),
	}
}
