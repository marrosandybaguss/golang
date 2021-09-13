package main

import "fmt"

func main() {
	// array key - value
	// jumlah data yg dimasukkan ke dlm Map boleh sebanyak-banyaknya

	person := map[string]string{ // [type key]type value
		"name": "Marro",
		"address": "Purwodadi",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["value"] = "TryToBeGoodPerson"
	fmt.Println(person["value"])

	values := make(map[string]int)
	values["lower"] = 55
	values["middle"] = 123
	values["higher"] = 212
	fmt.Println(values["lower"])
	fmt.Println(len(values))
	delete(values, "middle")
	fmt.Println(len(values))
}