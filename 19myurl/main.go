package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghp_1234567890"

func main() {
	fmt.Println(myurl)
	res, _ := url.Parse(myurl)
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Port())
	fmt.Println(res.Path)
	fmt.Println(res.RawQuery)

	//key value pairs

	qparams := res.Query()
	fmt.Printf("The type is :%T", qparams)

	fmt.Println("The value of coursename is :", qparams["coursename"])

	for key,val:=range qparams{
		fmt.Println(key,val)
	}


	chunksOfUrl:=&url.URL{ 
		Scheme: "https",
		Host:"codeforces.com",
		Path: "/profile",
		RawQuery: "MasterMuskan",
	}
	newUrl:=(chunksOfUrl).String()
	fmt.Println("The new url is :",newUrl)
}
