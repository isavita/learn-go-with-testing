package structs_methods_interfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Reactangle struct {
	Width  float64
	Height float64
}

func (r Reactangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Reactangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Perimeter() float64 {
	hypotenuse := t.calcHypotenuse()
	return hypotenuse + t.Base + t.Height
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) calcHypotenuse() float64 {
	return math.Sqrt(t.Base*t.Base + t.Height*t.Height)
}
