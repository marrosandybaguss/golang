package main

import "fmt"

func main() {
	fmt.Println("Ini String")
	// menghihtung jumlah karakter string 
	fmt.Println("Jumlah string: ", len("Ini String"))
	// mengambil byte char dari string
	fmt.Println("Byte character ke-0: ", "Ini String"[0])
}