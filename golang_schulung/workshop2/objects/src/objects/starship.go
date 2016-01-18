package objects

import ()

type Starship struct {
	Item
	Speed int
}

func NewStarship(name string, speed int) *Starship {
	ship := &Starship{
		Speed: speed,
	}
	ship.Item.Name = name
	return ship
}

func (ship *Starship) MoveInDirection(vector Point, time int) {
	for i := 0; i < ship.Speed; i++ {
		ship.Item.MoveInDirection(vector, time)
	}
}
