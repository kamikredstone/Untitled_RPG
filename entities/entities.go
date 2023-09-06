package entities

import (
	"github.com/kamikredstone/Untitled_RPG/items"
	"github.com/kamikredstone/Untitled_RPG/world"
)

type Entity interface {
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
