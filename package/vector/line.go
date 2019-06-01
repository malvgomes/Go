package main

import "math"

// Uppercase first letter = Public  (visible inside and outside the package)
// Lowercase first letter = Private (visible only inside the package)

// Point represents a coordinate in the plane
type Point struct {
	x float64
	y float64
}

func vector(a, b Point) (float64, float64) {
	cx := math.Abs(b.x - a.x)
	cy := math.Abs(b.y - a.y)
	return cx, cy
}

// Distance returns the distance between two points
func Distance(a, b Point) float64 {
	cx, cy := vector(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
