package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://example.com/sample?username=Aruna&age=20"

func main() {

	fmt.Println("Welcome to URLs")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	params := result.Query()
	// fmt.Printf("Type of parameter is: %T", params)
	// fmt.Println(params)
	for _, val := range params {
		fmt.Println(val)
	}
}
