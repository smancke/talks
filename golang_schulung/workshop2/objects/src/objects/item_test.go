package objects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Item_MoveTo(t *testing.T) {

	atom := NewItem("Atom")
	atom.MoveTo(Point{1, 1})
	atom.MoveTo(Point{2, 3})

	expected := Point{3, 4}
	if expected != atom.Pos() {
		t.Logf("Item not moved correct: expected=%v, is=%v", expected, atom.Pos())
		t.Fail()
	}
}

func Test_Item_MoveTo_WithAssert(t *testing.T) {
	atom := NewItem("Atom")
	atom.MoveTo(Point{1, 1})
	atom.MoveTo(Point{2, 3})

	a := assert.New(t)
	a.Equal(Point{3, 4}, atom.Pos())
}

func Test_Item_MoveInDirection(t *testing.T) {
	atom := NewItem("Atom")
	atom.MoveInDirection(Point{1, 2}, 5)

	a := assert.New(t)
	a.Equal(Point{5, 10}, atom.Pos())
}
