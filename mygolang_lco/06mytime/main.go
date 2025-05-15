package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcom to time Study of GoLang")

	presentTime := time.Now()
	// fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //Only use this format not any other

	createdDate := time.Date(2020, time.July, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
