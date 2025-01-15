package main

import (
	"Go-Quickstart-Caleb-Curry/points"
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Hello World in Go
	fmt.Println("Hello World!")

	// Variable declarations/assignments & types
	var poi int
	poi = 5
	fmt.Println(poi)
	poi2 := 45
	fmt.Println(poi2)
	poi2 = 50
	fmt.Println(poi2)

	FizzBuzz(15)
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

	fmt.Print("just enter your prize number : ")
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
	fmt.Print("\n")

	//error handling
	var numval int64
	for {
		fmt.Println("Please enter a valid number only")
		value, err := readinput.ReadString('\n')
		if err != nil {
			fmt.Println("Something went wrong with your input.")
			continue
		}

		value = strings.TrimSpace(value)
		numval, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			fmt.Println("You didn't enter a number, Try again.")
			continue
		}
		fmt.Printf("You successfully entered a num : %v\n", numval)
		break
	}
	// Understanding panics
	panix()

	//Maps
	message := `In an ancient world, where skies shimmered purple, a purple young girl named Elara lived.
				She discovered an ancient purple map, hidden in her purple attic, leading to an unknown treasure.
				Brimming with excitement, Elara embarked on her journey, accompanied by her loyal dog, Max.`
	PrintWordCount(CountWordsOccurence(message))

	//Structs
	point1 := points.Point{X: 7.4, Y: 3.4}
	point2 := points.Point{X: 10.2, Y: 9}
	pointslice := []points.Point{point1, point2}
	fmt.Println(pointslice[0].Y)

	fmt.Println("The bigger one is : ", point1.CompareTo(point2))
	fmt.Println(point1)

	collPoints := points.Magnitudes{
		points.Point{X: 8, Y: 6},
		points.Point{X: 2, Y: 1},
		points.Point{X: 9, Y: 5},
		points.Point{X: 1, Y: 2},
		points.Point{X: 3, Y: 1},
		points.Point{X: 5, Y: 4},
	}
	sort.Sort(collPoints)
	fmt.Println(collPoints)

	//interface & Polymorphism example
	pointco := points.Magnitudes{
		points.Point{X: 2, Y: 9},
		points.Point{X: 8, Y: 4},
		points.Point{X: 1, Y: 2},
		points.Point{X: 4, Y: 1},
		points.PointXYZ{X: 3, Y: 4, Z: 9},
		points.PointXYZ{X: 1, Y: 9, Z: 4},
		points.PointXYZ{X: 3, Y: 8, Z: 1},
	}
	sort.Sort(pointco)

	for _, value := range pointco {
		fmt.Printf("%s\t", value)
	}
	fmt.Println()

	// Basics of pointers
	class := 300
	ptr := &class
	fmt.Println(*ptr)
	*ptr = 400
	fmt.Println(class)

	xx, yy := 5, 10
	swap(&xx, &yy)
	fmt.Println("xx is : ", xx, "yy is : ", yy)

	valpoint := points.Point{X: 3, Y: 7}
	pntrpoint := &valpoint

	fmt.Println(pntrpoint)
	ModifyPoint(pntrpoint)
	fmt.Println(pntrpoint)

	intarr := [...]int{1, 2, 10}
	ModifyArray(&intarr)
	fmt.Println(intarr)
	ModifySlice(intarr[:])
	fmt.Println(intarr)

	//Goroutines
	channel := make(chan string)
	go Greet("Big Sula", channel)
	mssg := <-channel
	fmt.Println(mssg)
	time.Sleep(5 * time.Second)
}

//function Go-routine

func Greet(name string, channel chan string) {
	time.Sleep(3 * time.Second)
	channel <- fmt.Sprintln("hello", name)
}

func swap(x, y *int) {
	fmt.Println("Inside the function, before swap", *x, *y)
	*x, *y = *y, *x
	fmt.Println("Inside the function, after swap", *x, *y)
}

func ModifyPoint(p *points.Point) {
	p.X = 3000
}
func ModifyArray(arr *[3]int) {
	arr[0] = 5000
}
func ModifySlice(sli []int) {
	sli[0] = 233
}

//understanding panics

func panix() {
	// Understanding panics
	tstarr := [...]int{80, 90, 45}
	rand.Seed(time.Now().UnixNano())
	ari := rand.Intn(100) + 3
	fmt.Println("Before panic")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("A panic has occurred. This function recovers it")
		}
	}()
	fmt.Println(tstarr[ari])
}

// understanding functions

func PrintWordCount(wordcounts map[string]int) {
	for key, value := range wordcounts {
		fmt.Println(key, value)
	}
}

func CountWordsOccurence(message string) map[string]int {
	wordcounts := map[string]int{}
	messagewords := strings.Split(message, " ")
	for k := 0; k < len(messagewords); k++ {
		messagewords[k] = strings.ToLower(messagewords[k])
		messagewords[k] = strings.ReplaceAll(messagewords[k], " ", "")
		messagewords[k] = strings.ReplaceAll(messagewords[k], "\t", "")
		messagewords[k] = strings.ReplaceAll(messagewords[k], ",", "")
		messagewords[k] = strings.ReplaceAll(messagewords[k], ".", "")
	}
	for _, messageword := range messagewords {
		wordcounts[messageword]++
	}
	return wordcounts
}

// Fizzbuzz
func FizzBuzz(stop int) {
	for i := 1; i < stop; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz", " ")
			continue
		}
		if i%5 == 0 {
			fmt.Print("Buzz", " ")
			continue
		}
		if i%5 == 0 && i%3 == 0 {
			fmt.Print("FizzBuzz", " ")
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
}
