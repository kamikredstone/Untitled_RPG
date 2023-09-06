package entities

import (
	"github.com/kamikredstone/Untitled_RPG/items"
	"github.com/kamikredstone/Untitled_RPG/world"
)

type Entity interface {
	//Getters
	GetName() string
	GetLevel() *int
	GetStats() *CharacterStats
	GetInventory() string
	GetEquipment() string
	GetGraphic() string
	GetPosition() *EntityPosition
	GetHP() *int
	GetMP() *int
	GetInitiative() *int
	//Movement
	ChangePosition(EntityPosition) error
}

type EntityPosition struct {
	X int
	Y int
}

type CharacterStats struct {
	Strength     int
	Constitution int
	Dexterity    int
	Intelligence int
	Charisma     int
}

type PlayerCharacter struct {
	Name       string
	Level      *int
	HP         *int
	MP         *int
	Initiative *int
	Stats      *CharacterStats
	Inventory  []items.Item
	Equipment  []items.Item
	Graphic    string
	Position   *EntityPosition
	Room       *world.Room
}

func (p *PlayerCharacter) GetName() string {
	return p.Name
}

func (p *PlayerCharacter) GetStats() *CharacterStats {
	return p.Stats
}

func (p *PlayerCharacter) GetInventory() []items.Item {
	return p.Inventory
}

func (p *PlayerCharacter) GetEquipment() []items.Item {
	return p.Equipment
}

func (p *PlayerCharacter) GetPosition() *EntityPosition {
	return p.Position
}

func (p *PlayerCharacter) GetHP() *int {
	return p.HP
}

func (p *PlayerCharacter) GetMP() *int {
	return p.MP
}

func (p *PlayerCharacter) GetLevel() *int {
	return p.Level
}

func (p *PlayerCharacter) GetInitiative() *int {
	return p.Initiative
}

func (p *PlayerCharacter) ChangePosition(newPosition EntityPosition) {
	p.Position.X = newPosition.X
	p.Position.Y = newPosition.Y
}

func CreateDummyPlayerCharacter(room *world.Room) *PlayerCharacter {
	charName := "TestName"
	Level := 1
	HP := 10
	MP := 10
	Initiative := 0
	Stats := &CharacterStats{
		Strength:     1,
		Constitution: 1,
		Dexterity:    1,
		Intelligence: 1,
		Charisma:     1,
	}
	Inventory := make([]items.Item, 0)
	Equipment := make([]items.Item, 0)
	Graphic := "@"
	Position := EntityPosition{
		X: 0,
		Y: 0,
	}
	return &PlayerCharacter{
		Name:       charName,
		Level:      &Level,
		HP:         &HP,
		MP:         &MP,
		Initiative: &Initiative,
		Stats:      Stats,
		Inventory:  Inventory,
		Equipment:  Equipment,
		Graphic:    Graphic,
		Position:   &Position,
		Room:       room,
	}
}
