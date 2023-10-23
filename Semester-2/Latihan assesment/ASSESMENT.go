package main

import "fmt"

const NMAX int = 100

type countries struct {
	gdp      [NMAX]int
	nama     [NMAX]string
	nCountry int
}

func search(C countries, X string) int {
	var i int
	var idx int

	idx = NMAX
	fmt.Println("C countr = ", C.nCountry)
	for i < C.nCountry {

		if C.nama[i] == X {
			idx = i
			fmt.Println("INDEX = ", idx)
		} else {
			idx = NMAX
		}
		i++
	}
	return idx
}

func main() {
	var A countries

	entryData(&A)
	fmt.Println(A.nCountry)

	fmt.Println()
	// fmt.Println(A.nama[0], A.gdp[0])
	// fmt.Println(A.nama[1], A.gdp[1])
	// fmt.Println(A.nama[2], A.gdp[2])
	cetak(A)

}

func entryData(C *countries) {
	var name string
	var idx, gdp int
	var i int

	i = 0
	fmt.Scan(&name)
	for name != "NULL" {

		fmt.Scan(&gdp)
		C.nCountry = i
		idx = search(*C, name)

		if idx != NMAX {
			C.gdp[idx] = gdp // Mengganti nilai gdp di indeks idx dengan gdp baru
		} else {
			C.nama[i] = name
			C.gdp[i] = gdp
			i++
		}

		fmt.Scan(&name)

	}

	C.nCountry = i
}

func cetak(C countries) {
	var i int

	for i < C.nCountry {
		fmt.Println(C.nama[i], C.gdp[i])
		i++
	}

}
