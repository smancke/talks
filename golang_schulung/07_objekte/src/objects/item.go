package objects

import (
	"fmt"
)

type Point [2]int

func (p Point) Add(pointToAdd Point) Point {
	p[0] = p[0] + pointToAdd[0]
	p[1] = p[1] + pointToAdd[1]
	return p
}

type Item struct {
	Name string
	pos  Point
}

func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}

func (item *Item) Pos() Point {
	return item.pos
}

func (item *Item) MoveTo(vector Point) {
	item.pos = item.pos.Add(vector)
}

func (item *Item) MoveInDirection(vector Point, time int) {
	for i := 0; i < time; i++ {
		item.pos = item.pos.Add(vector)
	}
}

func (item *Item) String() string {
	return fmt.Sprintf("%v [%v,%v]", item.Name, item.pos[0], item.pos[1])
}
