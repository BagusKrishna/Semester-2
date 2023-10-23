package main

import "fmt"

func main() {
	var tlahir, totblanja, A, CD, diskon, dibayar int

	fmt.Scan(&tlahir, &totblanja)
	A = tlahir % 100
	CD = tlahir / 1000
	diskon = A * CD
	dibayar = totblanja - ((totblanja * diskon) / 100)

	fmt.Println("A", A)
	fmt.Println("CD", CD)
	fmt.Println("besar diskon:", diskon, "%")
	fmt.Println("Jumlah yang dibayar", dibayar)
}
