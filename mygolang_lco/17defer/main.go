package main

import "fmt"

func main() {

	defer fmt.Println("Defer Statement 1")
	defer fmt.Println("Defer Statement 2")
	fmt.Println("Welcome to Defer")
	myDefer()

}
func myDefer() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
