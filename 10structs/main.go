package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Active bool
}

func main() {
	Muskan := User{"Muskan", 23, "muskan@gmail.com", true}
	fmt.Println(Muskan)
	fmt.Printf("The user is : %+v\n",Muskan)
	fmt.Printf("The name is %v and age is %v",Muskan.Name,Muskan.Age)
}
