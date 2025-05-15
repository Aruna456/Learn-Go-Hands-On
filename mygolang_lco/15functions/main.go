package main

import "fmt"

func main() {

	fmt.Println("Welcome to functions!")

	// result :=
	fmt.Println("Result is:", add(1, 2, 3, 4, 5, 6))

}

func add(values ...int) int {

	tot := 0
	for _, i := range values {
		tot += i
	}
	return tot

}
