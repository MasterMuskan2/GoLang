package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(6) + 1

	fmt.Println("Value of number is ", num)

	switch num {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	case 3:
		fmt.Println("Number is 3")
	case 4:
		fmt.Println("Number is 4")
		fallthrough
	case 5:
		fmt.Println("Number is 5")
	case 6:
		fmt.Println("Number is 6")
	}
}
