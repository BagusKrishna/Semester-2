package main

import (
	"fmt"
)

func main() {
	var x1, y1, r1, x2, y2, r2, Xs, Ys int

	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&Xs, &Ys)

	inCircle1 := didalam(x1, y1, r1, Xs, Ys)
	inCircle2 := didalam(x2, y2, r2, Xs, Ys)

	if inCircle1 && inCircle2 {
		fmt.Println("Titik berada pada irisan lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Println("Titik berada pada lingkaran 1")
	} else if inCircle2 {
		fmt.Println("Titik berada pada lingkaran 2")
	} else {
		fmt.Println("Titik tidak berada pada kedua lingkaran atau irisan keduanya")
	}
}
func jarak(a, b, c, d int) int {
	return ((a - c) * (a - c)) + ((b - d) * (b - d))
}

func didalam(x, y, r, Xs, Ys int) bool {
	if jarak(x, y, Xs, Ys) <= r {
		return true
	} else {
		return false
	}
}

func akar(x float64) float64 {
	var a, b, c, d int
	x = float64(jarak(a, b, c, d))
	return x
}
