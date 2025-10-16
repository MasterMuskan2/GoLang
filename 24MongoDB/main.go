package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mastermuskan22/mongodbapi/router"
)

func main() {
	fmt.Println("MY API")
	r := router.Router()
	fmt.Println("The server is getting started.......")
	log.Fatal(http.ListenAndServe(":2000", r))
	fmt.Println("Listening to the port 2000.......")
}
