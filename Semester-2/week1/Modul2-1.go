package main

import (
	"fmt"
	"math"
)

func main() {
	var panjang, lebar int
	var luas, keliling int
	var diagonal float64

	fmt.Scan(&panjang, &lebar)

	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	diagonal = math.Sqrt(float64((panjang * panjang) + (lebar * lebar)))

	fmt.Println("Luas :", luas)
	fmt.Println("Keliling:", keliling)
	fmt.Println("Panjang diagonal:", diagonal)
}
