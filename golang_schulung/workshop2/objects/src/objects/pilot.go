package objects

type Flyable interface {
	MoveTo(Point)
}

func flyToMiddleOfUniverse(f Flyable) {
	f.MoveTo(Point{0, 0})
}
