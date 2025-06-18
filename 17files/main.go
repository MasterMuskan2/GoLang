package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//writing the content of the file
	content := "Hi, my name is Master Muskan and I am currently learning GoLang"

	//create the file
	file, err := os.Create("./myfile.txt")

	//check for error
	getNilError(err)

	//writing the content to the file
	length, err := io.WriteString(file, content)

	getNilError(err)

	fmt.Println("The length of the file is ", length)

	//lets close the file
	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	getNilError(err)
	fmt.Println("The data is ->", string(data))
}

func getNilError(err error) {
	if err != nil {
		panic(err)
	}
}
