package main

import "fmt"

func main() {
	// var value = (((2+6)%3)*4 - 2) / 3
	// var isEqual = (value >= 4)
	var kiri = false
	var kanan = true

	var kiriDanKanan = kiri && kanan
	fmt.Printf("kiri && kanan \t(%t) \n", kiriDanKanan)

	var kiriAtauKanan = kiri || kanan
	fmt.Printf("kiri || kanan \t(%t) \n", kiriAtauKanan)

	var kiriBalik = !kiri
	fmt.Printf("!kiri \t\t(%t) \n", kiriBalik)

	// fmt.Printf("nilai %d (%t) \n", value, isEqual)
}
