package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	hasil1 := z_1301223088(a, b)
	hasil2 := z_1301223088(b, a)

	fmt.Printf("%.3f %.3f", hasil1, hasil2)
}

func z_1301223088(x, y int) float64 {
	hasil := (3 * float64(x)) / (4 * float64(y))
	hasil = math.Sqrt((6 * float64(y) * hasil))
	return hasil
}
