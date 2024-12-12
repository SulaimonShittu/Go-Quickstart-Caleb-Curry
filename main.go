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
	fmt.Printf("%.20v\n", floating1)

	// understanding string & runes
	var firstChar rune = 's'
	fmt.Println(firstChar)
	fmt.Printf("%c\n", firstChar)

	var name string = "Sulaimon"
	var last string = "Shittu"
	name = name + " " + last
	fmt.Println(name)

	var nchar byte = name[0]
	fmt.Println(nchar)
	fmt.Printf("%c\n", nchar)

	// Delve into arrays & slices
	var car string = "m£rcedes"
	// converting a string to a slice gives it as a copy
	// so a modification to the slice doesn't affect the string data
	carSlices := []rune(car)
	fmt.Printf("%c\n", carSlices[1])

	names := [...]string{"Shitta", "Sulaimon", "Eniolorunmife"}
	namess := names[:]
	namess[0] = "Shittu"
	fmt.Println(namess)
	fmt.Println(names)

	namess = append(namess, "CK")
	fmt.Println(namess)
	fmt.Println(names)

	ages := []int{19, 21, 20, 22, 23, 18}
	for i, age := range ages {
		fmt.Println(i, age)
	}

}
