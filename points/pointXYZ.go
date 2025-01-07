package points

import (
	"fmt"
	"math"
)

// understanding structs & methods
type PointXYZ struct {
	X, Y, Z float64
}

func (p PointXYZ) Abs() float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2) + math.Pow(p.Z, 2))
}
func (po PointXYZ) String() string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f)", po.X, po.Y, po.Z)
}
