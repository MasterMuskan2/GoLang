package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("Value of the pointer is ", ptr)

	num:=22 
	fmt.Println("Value of the number is ", num)
	fmt.Println("Valur of the address is", &num)
	fmt.Println("Value of the pointer is ", *(&num))
}
