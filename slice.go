package main

import "fmt"

func main() {
	// adalah potongan dr array
	// ukuran slice bisa berubah beda dgn array
	// berkomunikasi dengan slice tp si slice ini nyimpen datanya dlm array 
	
	// pointer, length len(slice), dan capasisty cap(slice)
	var months = [...]string{
		"Januari",
		"Feb",
		"Maret",
		"Arpil",
		"Mei",
		"Juni",
		"Juli",
		"Agust",
		"September",
		"Okt",
		"Nov",
		"Des",
	}
	var slice = months[4:7]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	var slice1 = months[9:]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "Oktober"
	slice1[1] = "November"
	slice1[2] = "Desember"
	fmt.Println(months)
	fmt.Println()

	slice2 := months[10:]
	fmt.Println(slice2)
	slice3 := append(slice2, "MelebihiKapasitas")
	fmt.Println(slice3)
	slice3[1] = "BukanDesember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months) // Walaupun slice3 datanya diubah, array months tidak berubah datanya, 
	// karena slice3 merupakan slice baru yg ngerefer ke array baru (membuat array baru scr otomatis) yg dibuat ketika append melebihi kapasistas
	fmt.Println()

	// Create slice tanpa buat array dlu scr manual
	// array dibuat otomatis dan disembunyiin dibelakang
	newSlice := make([]string, 2, 5) // length 2, cap 5
	newSlice[0] = "Marro"
	newSlice[1] = "Sandy"
	fmt.Println(newSlice)

	// Copy slice
	// pastikan ukurannya sama, klo gk entar kepotong
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	fmt.Println()

	iniArray := [...]int{1,2,3}
	iniJugaArray := [3]int{1,2,3}
	iniSlice := []int{3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniJugaArray)
	fmt.Println(iniSlice)
}