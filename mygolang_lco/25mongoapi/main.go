package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aruna456/mongoapi/router"
)

func main() {
	fmt.Println("Welcome to MONGOAPI DEMO")
	r := router.Router()
	fmt.Println("Server is Getting Started...")
	log.Fatal(http.ListenAndServe(":7777", r))
	fmt.Println("Listening at port 7777...")
}
