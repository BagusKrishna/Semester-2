package main

import (
	"fmt"
)

func main() {
	var nominal int
	fmt.Print("Masukkan nominal uang: ")
	fmt.Scan(&nominal)

	denominasi(nominal)
}

func denominasi(nominal int) {
	var lembar10000, lembar2000, lembar1000 int

	// denominasi 10000
	lembar10000 = nominal / 10000
	nominal = nominal % 10000

	// denominasi 2000
	lembar2000 = nominal / 2000
	nominal = nominal % 2000

	// denominasi 1000
	lembar1000 = nominal / 1000

	fmt.Printf("Lembar 10000: %d\n", lembar10000)
	fmt.Printf("Lembar 2000: %d\n", lembar2000)
	fmt.Printf("Lembar 1000: %d\n", lembar1000)
}
