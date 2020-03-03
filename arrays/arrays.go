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

// PrintArrayLength should print array length.
func PrintArrayLength() {
	var names = [...]string{"Hello", "World"}
	fmt.Println(len(names))
}

// PrintArrayUsingIteration should print something.
func PrintArrayUsingIteration() {
	var numbers = [...]int{10, 11, 55, 12, 6}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("[%d] -> [%d]", i, numbers[i])
		fmt.Println()
	}
}

// PrintArrayUsingIteration2 should print something.
func PrintArrayUsingIteration2() {
	var programmingLanguages = [...]string{"Go", "Swift", "Kotlin", "Java", "PHP"}
	for index, programmingLanguage := range programmingLanguages {
		fmt.Println("[", index, "]", "->", programmingLanguage)
	}
}
