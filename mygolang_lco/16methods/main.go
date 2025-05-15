package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods!")

	Aruna := User{"Aruna", "aru@gmail", 19, true}
	fmt.Println(Aruna)
	Aruna.GetStatus()
	Aruna.SetEmail()
	fmt.Println(Aruna)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Status of the user is:", u.Status)
}

func (u User) SetEmail() {
	u.Email = "aruna@gmail.com"
	fmt.Println("New email:", u.Email)
}
