package entities

import (
	"errors"
)

type Entity interface {
	GetGraphic() string
	GetName() string
	GetPosition() map[string]int
	SetPosition(int, int) error
}

type Player struct {
	Name     string
	Graphic  string
	Position map[string]int
}

func NewPlayer(name string, graphic string, position map[string]int) *Player {
	return &Player{
		Name:     name,
		Graphic:  graphic,
		Position: position,
	}
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) GetGraphic() string {
	return p.Graphic
}

func (p *Player) GetPosition() map[string]int {
	return p.Position
}

func (p *Player) SetPosition(x int, y int) error {
	if p.Position["X"] == x && p.Position["Y"] == y {
		return errors.New("the position of the character hasn't changed")
	} else {
		p.Position["X"] = x
		p.Position["Y"] = y
	}
	return nil
}
