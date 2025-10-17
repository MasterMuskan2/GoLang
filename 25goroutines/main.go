package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //pointer

func main() {
	// go greeter("Hello")
	// greeter("World")
	websites := []string{
		"https://fb.com",
		"https://google.com",
		"https://github.com",
	}

	for _,web := range websites{
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string){

	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil{
		fmt.Println("We got some error!!!!")
	}else{
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
