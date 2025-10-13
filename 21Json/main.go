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
	// encodeJson()
	decodeJson()
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

func decodeJson(){
	data := []byte(`
		{
			"coursename": "dsa",
			"Price": 200,
			"Platform": "leetcode",
			"tags": ["dsa", "c++"]
        }
	`)

	var newCourse course

	isValid := json.Valid(data)

	if isValid{
		fmt.Println("The JSON was valid")
		json.Unmarshal(data, &newCourse)
		fmt.Printf("%#v\n", newCourse)
	}else{
		fmt.Println("The JSON was not valid")
	}

	var newData map[string]interface{}
	json.Unmarshal(data,&newData)
	fmt.Printf("%#v\n",newData)

	for key, value := range newData{
		fmt.Printf("The key is %v and value is %v and the type is %T\n", key, value, value)
	}
}