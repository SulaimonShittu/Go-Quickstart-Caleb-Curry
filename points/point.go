package points

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (po Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", po.X, po.Y)
}

func (p Point) Abs() float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
}
func (p Point) CompareTo(other Point) Point {
	if p.Abs() > other.Abs() {
		return p
	} else {
		return other
	}
}
