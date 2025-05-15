package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://example.com"

func main() {
	fmt.Println("Welcome to WEB REQUESTS")

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databyte))
}
