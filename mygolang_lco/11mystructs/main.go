package main

import "fmt"

func main() {

	fmt.Println("Structs in GoLang")

	Aruna := User{"Aruna", "arunasubramanian456@gmail.com", true, 19}
	fmt.Printf("Aruna details are %+v", Aruna)
	fmt.Println(Aruna.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
