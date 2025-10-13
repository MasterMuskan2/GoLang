package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World!")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serve).Methods("GET")
	log.Fatal(http.ListenAndServe(":2000", r))
}

func greeter() {
	fmt.Println("Hello there, I am Master Muskan")
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my Website</h1>"))
}
