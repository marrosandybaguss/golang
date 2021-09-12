package main

import "fmt"

func main() {
	const firstName string = "Marro Sandy"
	const lastName = "Bagus Saputra"
	const age = 37

	fmt.Println(firstName, lastName)
	fmt.Println("Age: ", age)

	const (
		firstName2 = "Sardoni"
		lastName2 = "Jagh"
	)


	fmt.Println(firstName2, lastName2)
	fmt.Println("Age: ", age)
}