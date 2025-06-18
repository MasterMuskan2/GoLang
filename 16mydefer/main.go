package main

import "fmt"

func main() {

	//defer works as LIFO

	defer fmt.Println("Master")
	defer fmt.Println("Muskan")
	defer fmt.Println("World")
	fmt.Println("Hello")
	printNum()
}

func printNum() {
	for i := range 5 {
		defer fmt.Println("Current value is ", i)
	}
}
