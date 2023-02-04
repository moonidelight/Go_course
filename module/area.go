package module

import (
	"math"
)

type Shape interface {
	area() float64
	length() float64
}

type Circle struct {
	R float64
}
type Rectangle struct {
	Width, Height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.R * circle.R
}

func (rec Rectangle) area() float64 {
	return rec.Width * rec.Height
}

func (circle Circle) length() float64 {
	return 2 * math.Pi * circle.R
}

func (rec Rectangle) length() float64 {
	return 2 * (rec.Width + rec.Height)
}
func Area(shape Shape) float64 {
	return shape.area()
}
func Length(shape Shape) float64 {
	return shape.length()
}
