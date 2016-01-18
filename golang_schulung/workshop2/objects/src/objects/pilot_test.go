package objects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_flyToMiddleOfUniverse_WithoutMock(t *testing.T) {
	herzAusGold := NewStarship("HerzAusGold", 42)

	flyToMiddleOfUniverse(herzAusGold)

	assert.Equal(t, Point{0, 0}, herzAusGold.Pos())
}

func Test_flyToMiddleOfUniverse_Mock(t *testing.T) {
	//.. TODO ..
}
