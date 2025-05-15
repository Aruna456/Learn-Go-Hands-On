package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Swtich case")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	switch diceNumber {
	case 1:
		fmt.Println("You can move a coin 1 time!")
	case 2:
		fmt.Println("You can move 2 times!")
	case 3:
		fmt.Println("You can move 3 times!")
	case 4:
		fmt.Println("You can move 4 times!")
	case 5:
		fmt.Println("You can move 5 times!")
	case 6:
		fmt.Println("You can move 6 times and get another chance!")
	default:
		fmt.Println("Try again!")
	}
}
