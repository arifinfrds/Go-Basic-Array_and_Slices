package arrays

import "fmt"

// PrintDefaultArray should do something.
func PrintDefaultArray() {
	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a)
}

// PrintCustomArray should do something.
func PrintCustomArray() {
	var numbers = [...]int{14, 5, 123, 33}
	fmt.Println(numbers)
}

// PrintStringArray should do something.
func PrintStringArray() {
	var names = [...]string{"Hello", "World"}
	fmt.Println(names)
	fmt.Println(names[0])
}
