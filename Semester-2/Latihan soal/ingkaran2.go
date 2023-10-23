package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	posisi := posisiTitik(cx1, cy1, r1, cx2, cy2, r2, x, y)
	fmt.Println(posisi)
}

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func posisiTitik(cx1, cy1, r1, cx2, cy2, r2, x, y float64) string {
	diLingkaran1 := didalam(cx1, cy1, r1, x, y)
	diLingkaran2 := didalam(cx2, cy2, r2, x, y)

	if diLingkaran1 && diLingkaran2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if diLingkaran1 {
		return "Titik di dalam lingkaran 1"
	} else if diLingkaran2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}
