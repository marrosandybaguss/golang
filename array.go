package main

import "fmt"

func main() {
	// type data dlm array hrs sama
	// hrs ditentukan jumlah data yg bisa ditampung
	// daya tampung tdk dapat ditambah
	var names [3]string
	names[0] = "Marrosandy"
	names[1] = "Bagus"
	names[2] = "Saputra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int {
		81,
		87,
		93,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// panjang array, tp bukan jumlah data didalam array 
	var valNumber = len(values)
	fmt.Println(valNumber)

}