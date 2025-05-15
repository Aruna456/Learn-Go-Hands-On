package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in GoLang!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Water Melon"
	// fruitList[0]="Apple"
	fmt.Println("FruitList is", fruitList)
	fmt.Println("FruitList is", len(fruitList))

	var veggieList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("VeggieList is", veggieList)

}
