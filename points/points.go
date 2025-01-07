package points

import "fmt"

// interfaces cont
type Points []Point

func (points Points) Len() int {
	return len(points)
}
func (points Points) Less(i, j int) bool {
	return points[i].Abs() < points[j].Abs()
}
func (points Points) Swap(i, j int) {
	points[i], points[j] = points[j], points[i]
}
func (points Points) String() string {
	str := ""
	for _, i2 := range points {
		str += i2.String() + " : " + fmt.Sprintf("%f", i2.Abs()) + "\n"
	}
	return str
}
