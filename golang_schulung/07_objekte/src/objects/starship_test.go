package objects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Starship_MoveTo(t *testing.T) {
	herzAusGold := NewStarship("HerzAusGold", 42)
	herzAusGold.MoveTo(Point{1, 1})
	herzAusGold.MoveTo(Point{2, 3})

	a := assert.New(t)
	a.Equal(Point{3, 4}, herzAusGold.Pos())
}

func Test_Starship_MoveInDirection(t *testing.T) {
	herzAusGold := NewStarship("HerzAusGold mit Unwahrscheinlichkeitsdrive", 42)
	herzAusGold.MoveInDirection(Point{1, 2}, 5)

	a := assert.New(t)
	a.Equal(Point{210, 420}, herzAusGold.Pos())
}
