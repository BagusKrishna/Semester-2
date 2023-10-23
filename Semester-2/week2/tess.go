package main

import "fmt"

func isTriangle(a, b, c int) bool {
	// Mengecek apakah tiga sisi a, b, c dapat membentuk segitiga.
	return a+b > c && a+c > b && b+c > a
}

func main() {
	var a, b, c int
	fmt.Println("Masukkan nilai sisi bidang: ")
	fmt.Scanln(&a, &b, &c)

	if isTriangle(a, b, c) {
		fmt.Println("segitiga")
	} else {
		fmt.Println("bukan segitiga")
	}
}
