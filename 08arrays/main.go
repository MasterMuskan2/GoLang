package main

import (
	"fmt"
	"sort"
)

func main() {
	var names = []string{"John", "Jane", "Jim"}
	fmt.Println("Length ", len(names))
	fmt.Println(names)
	names = append(names, "Jerry")
	fmt.Println(names)
	names=append(names[1:len(names)])
	fmt.Println(names)


	age:=make([]int,5)
	age[0]=2
	age[1]=1
	age[2]=8
	age[3]=3
	age[4]=4
	fmt.Println(age)

	age = append(age, 5,11,0)
	fmt.Println(age)
	fmt.Println(sort.IntsAreSorted(age))
	sort.Ints(age)
	fmt.Println(age)
}
