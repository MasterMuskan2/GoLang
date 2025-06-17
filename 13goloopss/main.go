package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesady", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("The index is %v and the day is %v\n", index, day)
	}

	i := 0

	for i <= 10 {

		if i == 3 {
			break
		}

		if i == 4 {
			i++
			continue
		}

		if i == 5 {
			goto hii
		}

		fmt.Println("The value is ", i)
		i++
	}
hii:
	fmt.Println("Master Muskan")
}
