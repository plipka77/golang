package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Volume() float64
	Kind() string
}

type ShapeSlice []Shape

func (s ShapeSlice) Len() int {
	return len(s)
}

func (s ShapeSlice) Less(i, j int) bool {
	return s[i].Volume() < s[j].Volume()
}

func (s ShapeSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ShapeSlice) Print() {
	for i, shape := range s {
		fmt.Printf("Order: %d  Shape: %s  Volume: %.2f\n", i, shape.Kind(), shape.Volume())
	}
}

type Sphere struct {
	Radius float64
}

func (s Sphere) Volume() float64 {
	return (4 / 3) * math.Pi * math.Pow(s.Radius, 3)
}

func (s Sphere) Kind() string {
	return "Sphere"
}

type Cube struct {
	Edge float64
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Edge, 3)
}

func (c Cube) Kind() string {
	return "Cube"
}

type Cylinder struct {
	Radius float64
	Height float64
}

func (c Cylinder) Volume() float64 {
	return math.Pi * math.Pow(c.Radius, 2) * c.Height
}

func (c Cylinder) Kind() string {
	return "Cylinder"
}
