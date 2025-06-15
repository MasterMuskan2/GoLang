package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "John"
	names[1] = "Jane"
	// names[2] = "Jim"
	fmt.Println(names)
	fmt.Println("Length of the array is ", len(names))

	var age=[4] int{1,2,3,4}
	fmt.Println(age)
}
