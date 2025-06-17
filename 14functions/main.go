package main

import "fmt"

func getHello() {
	fmt.Println("Hii, my name is Master Muskan")
}

func getSum(val1 int, val2 int) int {
	return (val1+val2)
}

func getMoreSum(val ...int)int{
	sum:=0
	for i:=range(val){
		sum+=i
	}
	return sum
}

func main() {
	fmt.Println("Getting started here")
	getHello()
	res:=getSum(1,7)
	fmt.Println("Sum is ",res)
	ans:=getMoreSum(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("Big Sum is ",ans)
}
