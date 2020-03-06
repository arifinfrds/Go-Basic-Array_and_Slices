package slices

import "fmt"

// PrintSlicesFromArray should print slices
func PrintSlicesFromArray() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)

	var slices []int = numbers[0:3]
	fmt.Println("Slice before:")
	fmt.Println(slices)

	modifySlices(slices)
	modifySlices2(slices)

	fmt.Println("Slice after:")
	fmt.Println(slices)

	fmt.Println("Slice Cap:")
	printSlicesCapacity(slices)

	var newSlices = makeNewSlices(20)
	fmt.Println("NewSlices Cap:")
	fmt.Println(newSlices)

	fmt.Println("Append Slices")
	appendSlices()

	fmt.Println("appendSlicesWithNilChecking")
	appendSlicesWithNilChecking()

	fmt.Println("append food")
	appendSlicesFood()
}

func modifySlices(slices []int) {
	for i := 0; i < len(slices); i++ {
		slices[i]++
	}
}

func modifySlices2(slices []int) {
	for index := range slices {
		slices[index]++
	}
}

func printSlicesCapacity(slices []int) {
	fmt.Println(cap(slices))
}

func makeNewSlices(size int) []int {
	var newSlices = make([]int, size)
	return newSlices
}

func appendSlices() {
	var cars = []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars", cars, "has old length:", len(cars), "and capacity:", cap(cars))

	cars = append(cars, "Toyota")
	fmt.Println("cars", cars, "has old length:", len(cars), "and capacity:", cap(cars))

	cars = append(cars, "BMW")
	fmt.Println("cars", cars, "has old length:", len(cars), "and capacity:", cap(cars))
}

func appendSlicesWithNilChecking() {
	var names []string
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "Arifin", "Firdaus")
		fmt.Println("names contents:", names)
	}
}

func appendSlicesFood() {
	var veggies = []string{"potatoes", "tomatoes", "brinjal"}
	var fruits = []string{"oranges", "apples"}
	var food = append(veggies, fruits...)
	fmt.Println("food:", food)
}
