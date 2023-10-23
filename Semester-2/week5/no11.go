package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)
	hasil1 := pertama(a, b)
	hasil2 := pertama(b, a)
	fmt.Printf("%.3f %.3f", hasil1, hasil2)
}

func pertama(a, b int) float64 {
	z := (6 * float64(b)) * (3 * float64(a)) / (4.0 * float64(b))
	hasil := math.Sqrt(float64(z))
	return hasil
}
