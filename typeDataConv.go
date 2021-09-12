package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	fmt.Println()

	var name = "Marro Sandy"
	var r byte = name[2]
	var rString string = string(r)

	fmt.Println(name)
	fmt.Println(rString)
}