package main

import (
	"fmt"
)

// var jwtToken = 300000

const LoginToken string = "qwerty" //Public const

func main() {
	var username string = "Aruna"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)
	// fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloatVal float64 = 255.056342567598657
	fmt.Println(smallFloatVal)
	fmt.Printf("variable is of type: %T \n", smallFloatVal)

	//Default values and some aliases
	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "google.com"
	fmt.Println(website)

	//no var style
	numOfUser := "Aruna"
	fmt.Println(numOfUser)
	fmt.Printf("variable is of type: %T \n", numOfUser)

	// fmt.Println(LoginToken)
	// fmt.Printf("variable is of type: %T \n", LoginToken)

}
