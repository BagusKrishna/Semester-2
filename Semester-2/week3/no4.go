package main

import "fmt"

func main() {
	var tabungan, jumlah int
	var i, tertinggi, terendah int
	var rata_rata float64

	tertinggi = 0
	fmt.Scan(&tabungan)
	if tabungan >= 0 {
		terendah = tabungan
		for tabungan >= 0 {

			i++
			jumlah = jumlah + tabungan
			if tabungan > tertinggi {
				tertinggi = tabungan
			}
			if tabungan < terendah {
				terendah = tabungan
			}
			fmt.Scan(&tabungan)
		}
	}
	rata_rata = float64(jumlah) / float64(i)
	fmt.Println("Jumlah =", jumlah)
	fmt.Println("Rata - rata =", rata_rata)
	fmt.Println("Uang tertinggi =", tertinggi)
	fmt.Println("Uang terendah =", terendah)
}
