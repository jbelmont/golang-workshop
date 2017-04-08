package workshop

import (
	"math"
)

// Shape interface has an area, circumference, and volume methods
type Shape interface {
	area()
	circumference()
	volume()
}

// Circle is a type of shape that has a radius and uses the value PI
type Circle struct {
	radius float64
	PI     float64
}

func (circle *Circle) area() float64 {
	return circle.PI * math.Pow(circle.radius, 2)
}

func (circle *Circle) circumference() float64 {
	return 2 * circle.PI * circle.radius
}

// Square is a shape that has 4 equal sides
type Square struct {
	side float64
}

func (square *Square) area() float64 {
	return math.Pow(square.side, 2)
}

type Sphere struct {
	PI     float64
	radius float64
}

func (sphere *Sphere) volume() float64 {
	return (4 / 3) * sphere.PI * math.Pow(sphere.radius, 3)
}

// Cube is a 3-dimensional circle
type Cube struct {
	side float64
}

func (cube *Cube) volume() float64 {
	return math.Pow(cube.side, 3)
}

func (square *Square) volume() float64 {
	return math.Pow(square.side, 3)
}

// Rectangle is a 4 sided polygon
type Rectangle struct {
	height int
	width  int
}

func (rectangle *Rectangle) area() int {
	return rectangle.height * rectangle.width
}

func interfaces() {
	circle := Circle{radius: 5.5, PI: math.Pi}
	areaOfCircle := circle.area()

	assert(areaOfCircle == 55)
	assert(circle.circumference() == 88)

	var square = Square{5}
	assert(square.area() == 36)

	cube := new(Cube)
	assert(cube.volume() == 8)

	var rectangle = Rectangle{width: 4, height: 5}
	assert(rectangle.area() == 2)
}
