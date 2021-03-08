package shape

// Shape interface definition
type Shape interface {
	Area() float64
	Perimeter() float64
	Info() string
}

// shape struct definition
type shape struct {
	name      string
	shapeType string
}

// NewName sets a new name to shape s
func (s *shape) NewName(n string) {
	s.name = n
}

// NewShapeType sets a new FigureType to shape s
func (s *shape) NewShapeType(st string) {
	s.shapeType = st
}
