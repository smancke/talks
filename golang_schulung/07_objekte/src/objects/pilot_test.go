package objects

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:generate go get github.com/golang/mock/mockgen
//go:generate $GOPATH/bin/mockgen -self_package objects -package objects -destination $GOPATH/src/objects/mocks_test.go objects Flyable
func Test_flyToMiddleOfUniverse_WithoutMock(t *testing.T) {
	herzAusGold := NewStarship("HerzAusGold", 42)

	flyToMiddleOfUniverse(herzAusGold)

	assert.Equal(t, Point{0, 0}, herzAusGold.Pos())
}

func Test_flyToMiddleOfUniverse_Mock(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockFlyable(ctrl)

	//mock.EXPECT().MoveTo(Point{1, 0})
	mock.EXPECT().MoveTo(gomock.Any()).Do(func(p Point) {
		assert.Equal(t, Point{0, 0}, p)
	})

	flyToMiddleOfUniverse(mock)
}

func Test_flyToMiddleOfUniverse_FunctionMock(t *testing.T) {

	mockWasCalled := false
	mock := func(p Point) {
		mockWasCalled = true
	}
	flyToMiddleOfUniverse(myFuncMock(mock))
	assert.True(t, mockWasCalled)

}

type myFuncMock func(p Point)

func (f myFuncMock) MoveTo(p Point) {
	f(p)
}
