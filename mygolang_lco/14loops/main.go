package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to loops!!")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }
	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto label
		}

		fmt.Println("Rougue value is", rougueValue)
		rougueValue++
	}

label:
	fmt.Println("Jumping at Loop")
}
