package main

import "fmt"

func main() {
	val := 10
	if val > 100 {
		fmt.Println("Master")
	} else if val < 100 {
		fmt.Println("Muskan")
	} else {
		fmt.Println("Master Muskan")
	}

	if age := 20; age >= 18 {
		fmt.Println("Hiii")
	} else {
		fmt.Println("Hello")
	}
}
