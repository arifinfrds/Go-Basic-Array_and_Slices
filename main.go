package main

import (
	"arrays_and_slices/arrays"
	"arrays_and_slices/slices"
	"fmt"
)

func main() {
	printArrays()
	printSlices()
}

func printArrays() {
	fmt.Println("printArrays()")
	arrays.PrintDefaultArray()
	arrays.PrintCustomArray()
	arrays.PrintStringArray()
	arrays.PrintArrayLength()
	arrays.PrintArrayUsingIteration()
	arrays.PrintArrayUsingIteration2()
}

func printSlices() {
	fmt.Println("printSlices()")
	slices.PrintSlicesFromArray()
}
