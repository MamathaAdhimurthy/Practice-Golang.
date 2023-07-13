package coding

import "math"

type Shape interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Traingle struct {
	Base   float64
	Height float64
	Side1  float64
	Side2  float64
	Side3  float64
}

func (t Traingle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Traingle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Side3
}
