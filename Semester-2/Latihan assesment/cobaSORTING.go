package main

import "fmt"

const NMAX = 100

type data struct {
	tabInt [NMAX]int
	n      int
}

func isi(T *data) {
	var i int

	fmt.Print("Banyak data : ")
	fmt.Scan(&T.n)
	fmt.Println("Masukkan data : ")
	for i < T.n {
		fmt.Scan(&T.tabInt[i])
		i++
	}
}

func cetak(T data) {
	var i int

	fmt.Println(T.n)
	for i < T.n {
		fmt.Print(T.tabInt[i], " ")
		i++
	}
}

// func sort(T data) {
// 	var i, j int
// 	var temp, idx int

// 	for j < T.n {
// 		idx = j - 1
// 		i = j
// 		for i < T.n {
// 			if T.tabInt[idx] < T.tabInt[i] {
// 				idx = i
// 			}
// 			i++
// 		}
// 		temp = T.tabInt[j-1]
// 		T.tabInt[j-1] = T.tabInt[idx]
// 		T.tabInt[idx] = temp
// 		j++
// 	}

// }

// func sortinser(T *data) {
// 	var pass, i, temp int

// 	pass = 0
// 	for pass < T.n {
// 		i = pass
// 		temp = T.tabInt[pass]
// 		for i > 0 && temp > T.tabInt[i-1] {
// 			T.tabInt[i] = T.tabInt[i-1]
// 			i--
// 		}
// 		T.tabInt[i] = temp
// 		pass++
// 	}
// }

// func sortinser(T *data) {
// 	var i, j int
// 	var temp int

// 	for i < T.n {
// 		j = i
// 		temp = T.tabInt[j]
// 		for j > 0 && temp > T.tabInt[j-1] {
// 			T.tabInt[j] = T.tabInt[j-1]
// 		}
// 		T.tabInt[j] = temp
// 	}
// }

func sortinser(T *data) {
	var i, j int
	var temp int

	for i < T.n {
		j = i
		temp = T.tabInt[j]
		for j > 0 && temp > T.tabInt[j-1] {
			T.tabInt[j] = T.tabInt[j-1]
			j--
		}
		T.tabInt[j] = temp
		i++
	}
}

func main() {
	var T data

	isi(&T)
	fmt.Println()
	fmt.Println("Data kamu :")
	cetak(T)
	sortinser(&T)
	fmt.Println()
	fmt.Println("Setelah di sort :")
	cetak(T)
}
