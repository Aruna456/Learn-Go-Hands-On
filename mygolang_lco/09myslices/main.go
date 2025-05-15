package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")
	var fruitList = []string{"Apple", "Tomato"}
	// fmt.Println("%T", fruitList)
	fruitList = append(fruitList, "Orange", "Watermelon", "Pineapple")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[:3]) //Slices the slice
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 400
	highScore[1] = 950
	highScore[2] = 200
	highScore[3] = 543

	highScore = append(highScore, 555, 666, 890)
	fmt.Println(highScore)
	// highScore[4] = 600
	sort.Ints(highScore)
	fmt.Println(highScore)

	var courses = []string{"ReactJs", "JS", "Swift", "Python", "Java", "Ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
