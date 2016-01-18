package objects

type flyable interface {
	MoveTo(Point)
}

func flyToMiddleOfUniverse(f flyable) {
	f.MoveTo(Point{0, 0})
}
