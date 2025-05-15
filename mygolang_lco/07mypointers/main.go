package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Pointers!")

	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value: ", myNumber)

}
