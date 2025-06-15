package main

import "fmt"

func main() {
	count := make(map[string]int)
	count["apple"] = 10
	count["banana"] = 20
	count["orange"] = 30
	fmt.Println(count)
	fmt.Println("Value of apple is ", count["apple"]) 
	fmt.Println("Value of mango is ", count["mango"]) 
	delete(count,"mango")
	fmt.Println(count)
	delete(count,"apple")
	fmt.Println(count)

	for key,value:=range count{
		fmt.Println("Key is",key, "Value is",value)
	}
}
