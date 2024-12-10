package main

import (
	"fmt"
	"math"
)

func main() {
	// Hello World in Go
	fmt.Println("Hello World!")

	// Variable declarations/assignments & types
	var points int
	points = 5
	fmt.Println(points)
	points2 := 45
	fmt.Println(points2)
	points2 = 50
	fmt.Println(points2)

	const age = 20
	fmt.Println(age)

	data := [5]int{2, 4, 5}
	fmt.Println(data)

	// Numeric types in Go

	maxint := math.MaxUint8
	fmt.Println(maxint)

	var integer1 int32 = 76
	var integer2 int64 = 99

	fmt.Println(int64(integer1) + integer2)

	var floating1 float64 = 56.173
	fmt.Printf("%.20v", floating1)

}
