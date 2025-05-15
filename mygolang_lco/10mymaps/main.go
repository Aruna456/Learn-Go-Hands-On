package main

import "fmt"

func main() {
	fmt.Println("Maps!!")

	languages := make(map[string]string)

	languages["Js"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("Map of Languages:", languages)
	fmt.Println("Js shorts for", languages["Js"])

	delete(languages, "RB")
	fmt.Println("Map of Languages:", languages)

	//loop through map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
