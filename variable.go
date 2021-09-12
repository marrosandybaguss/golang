package main

import "fmt"

func main() {
	var name string

	name = "Marro Sandy"
	fmt.Println(name)

	var age = 30
	fmt.Println("Age: ", age)

	name2 := "Bagus Saputra"
	fmt.Println(name2)

	var age2 int8 = 37
	fmt.Println("Age: ", age2)

	var (
		firstName = "Marro Sandy"
		lastName = "Bagus Saputra"
	)

	fmt.Println("My Name: ", firstName, lastName)

}