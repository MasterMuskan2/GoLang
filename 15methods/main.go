package main

import "fmt"

func main() {
	Muskan := User{"Muskan", 23, true, "muskan@gmail.com"}
	fmt.Println(Muskan)
	Muskan.getStatus()
	Muskan.changeMail()
	fmt.Println("The original email is ", Muskan.Email)
}

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

func (u User) getStatus() {
	fmt.Println("The status is ", u.Status)
}

func (u User) changeMail() {
	u.Email = "master@gmail.com"
	fmt.Println("The email is ", u.Email)
}
