package main

import "fmt"

const phi float64 = 3.14

func hitungVolume(r, t float64) float64 {
	volume := phi * r * r * t
	return volume
}

func hitungMassa(r, t, p float64) float64 {
	var massa float64
	massa = hitungVolume(r, t) * p
	return massa
}

func display(m1, m2 float64) {
	var selisih float64

	if m1 > m2 {
		selisih = m1 - m2
		fmt.Println("m1 lebih berat daripada m2 dengan selisih ", selisih)
	} else if m2 > m1 {
		selisih = m2 - m1
		fmt.Println("m2 lebih berat daripada m1 dengan selisih ", selisih)
	} else {
		fmt.Println("BALANCE")
	}

}

func main() {
	var r float64
	var tKanan, pKanan float64
	var tKiri, pKiri float64
	var kiri, kanan float64

	fmt.Scan(&r)
	fmt.Scan(&tKiri, &pKiri)
	fmt.Scan(&tKanan, &pKanan)

	kiri = hitungMassa(r, tKiri, pKiri)
	kanan = hitungMassa(r, tKanan, pKanan)

	display(kiri, kanan)
}
