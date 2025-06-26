package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// getRequest()
	postRequest()
}

func getRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code :", response.StatusCode)
	fmt.Println("Content length is :", response.ContentLength)
	content,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	fmt.Println("The content is :", string(content))
}

func postRequest(){
	const myurl="http://localhost:8000/post"
	requestBody:=strings.NewReader(`
	{
		"name":"Muskan",
		"age":20,
		"role":"software engineer"
	}
	`)
	response,err:=http.Post(myurl,"application/json",requestBody)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	fmt.Println("The content is :", string(content))
}
