package main

import (
	"sort"
)

func main() {
	cylinder := Cylinder{5, 10}
	sphere := Sphere{10}
	cube := Cube{5}

	shapeSlice := ShapeSlice{cylinder, sphere, cube}
	sort.Sort(shapeSlice)

	shapeSlice.Print()
}
