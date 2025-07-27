package structsmethodsinterfaces

import "math"

type Shape interface {
	// Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)

}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Height * t.Base
}
