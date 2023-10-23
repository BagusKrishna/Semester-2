package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	z1 := z(a, b)
	z2 := z(b, a)

	fmt.Printf("%.3f %.3f", z1, z2)
}

func z(x, y int) float64 {
	return math.Sqrt(float64(18 * x / (16 * y)))
}
