package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Hello World!")
	encodeJson()
}

func encodeJson() {
	newCourse := []course{
		{"web-dev", 100, "gfg", "abc123", []string{"web-dev", "js", "html"}},
		{"dsa", 200, "leetcode", "mno987", []string{"dsa", "c++"}},
		{"android", 500, "udemy", "pqr567", nil},
	}

	//package data as json
	// resultJson, err := json.Marshal(newCourse)
	resultJson, err := json.MarshalIndent(newCourse, "", "\t") //Indented based on "\t"
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", resultJson)
}
