package main

import "fmt"

func main() {
	var g string
	var jr, jl, totalA, totalB, totalSum int

	for g != "Z" {
		fmt.Scan(&g, &jr, &jl)
		if g == "A" {
			jr = jr * 75000
			jl = jl * 90000
			totalA = totalA + jr + jl
			fmt.Println(totalA)

		} else if g == "B" {
			jr = jr * 125000
			jl = jl * 70000
			totalB = totalB + jr + jl
			fmt.Println(totalB)
		} else if g == "Z" {
			fmt.Println("Biaya yang dikeluarkan perusahaan untuk gaji karyawan", totalSum)
		} else {
			fmt.Println("Golongan tidak dikenali")
		}
		totalSum = totalA + totalB

	}

}
