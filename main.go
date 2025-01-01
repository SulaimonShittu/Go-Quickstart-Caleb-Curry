package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	fmt.Printf("This Type : %v is a %T type\n", data, data)

	// Numeric types in Go
	maxint := math.MaxUint8
	fmt.Println(maxint)
	fmt.Println(math.MaxInt64 == math.MaxInt)

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

	//Essential operations & functions
	// Dive into operators
	var x int = 5
	var y float64 = 10

	fmt.Println(float64(x) / y)
	x++
	fmt.Println(x)

	scores := [...]int{7, 9, 10, 11, 12, 13}
	for _, val := range scores {
		val += 5
	}
	fmt.Println(scores)

	//important numeric functions
	fmt.Println(math.Sqrt(10))
	fmt.Println(math.Pow(10, 2))
	fmt.Println(math.Max(5, 10), math.Min(5, 10))
	fmt.Println(math.Ceil(1.1))
	fmt.Println(math.Abs(-10))

	fmt.Println(math.Sqrt(math.Inf(0)))

	pie := "3.142"
	pieV, _ := strconv.ParseFloat(pie, 64)
	fmt.Printf("The value: %f is of type: %T", pieV, pieV)

	// important string functions
	fmt.Println(strings.Contains("Sulaimon", "lai"))
	fmt.Println(strings.HasPrefix("Sulaimon", "S"))
	fmt.Println(strings.HasSuffix("sul mon", "mon"))
	fmt.Println(strings.Split("Oye Olohun Gbo Gbo Kan Ma Lo Ye olohun", " "))
	fmt.Println(strings.Join([]string{"Just", "Do", "It"}, " "))

	var sbv strings.Builder
	for i := 0; i < 4; i++ {
		sbv.WriteString("Iren")
	}
	sbvs := sbv.String()
	fmt.Println(sbvs)

	// important array and slice methods
	arrdata := [...]int{80, 78, 98, 79, 81, 86, 70, 90}
	fmt.Println(len(arrdata) == cap(arrdata))
	slicedata := arrdata[:]
	slicedata = append(slicedata, 65)
	fmt.Println(len(slicedata), cap(slicedata))

	slida := make([]int, 20, 50)
	fmt.Println(slida, len(slida), cap(slida))

	slida = arrdata[2:6]
	fmt.Println(slida)

	grades := []int{5, 10, 15, 20, 25, 30}
	grades2 := [10]int{}
	fmt.Println(copy(grades2[:], grades), grades, grades2)

	grades3 := make([]int, len(grades))
	copy(grades3, grades)
	grades[2] = 6
	fmt.Println(grades, grades3)

	//Essential of dynamic apps
	//Delve into constants and iota

	const bonus = 10
	var pay float64 = bonus + 700
	fmt.Println(pay)
	const years int = 20
	var due float64 = float64(years)
	println(due)

	const (
		gold = iota
		silver
		bronze
	)
	prizes := [...]string{"gold", "silver", "bronze"}

	// collecting user input

	readinput := bufio.NewReader(os.Stdin)
	fmt.Print("please input the prize value : 0, 1, 2 : ")
	input, _ := readinput.ReadString('\n')
	input = strings.TrimSpace(input)
	mbr, _ := strconv.ParseInt(input, 10, 64)
	println(mbr, prizes[mbr])

	var mbr2 int
	fmt.Scan(&mbr2)

	//if conditionals
	if mbr2 == 0 {
		fmt.Println("You are a one of one talent. congrats!")
	} else if mbr2 == 1 {
		fmt.Println("You are really good. you just fell short from the best.")
	} else if mbr2 == 2 {
		fmt.Println("good finish. You can build up on this next time")
	}

	// switch statements
	const (
		red = iota + 1
		green
		blue
	)
	var val int = 2

	switch val {
	case 0, 1:
		fmt.Println("Valid value but not optimal")
	case 2:
		fmt.Println("Optimal value")
	default:
		fmt.Println("Invalid value")
	}

	switch {
	case red < green && red < blue:
		fmt.Println("Red is the champion!")
	}

	//loops
	iterv := 1
	for iterv < 5 {
		fmt.Printf("The square of the number : %v is %v\n", iterv, math.Pow(float64(iterv), 2))
		iterv++
	}
	for i := 0; i <= 2; i++ {
		fmt.Print(i, "\t")
	}
}
