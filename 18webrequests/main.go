package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://leetcode.com/"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(data)
	fmt.Println("The content is -> ", content)
}
