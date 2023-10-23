package main

import "fmt"

const MAXSIZE = 999999

type Trafik [MAXSIZE]int
type urutan [MAXSIZE]int

func main() {
	var arr Trafik
	var n int
	var i int = 1

	catatLaluLintas(&arr, &n)
	fmt.Println("Data kamu :")
	for i <= n {
		fmt.Print(arr[i], " ")
		i++
	}

	fmt.Println()
	fmt.Println("Data kamu setelah diurutkan : ")
	urut(&arr, n)
	i = 1
	for i <= n {
		fmt.Print(arr[i], " ")
		i++
	}

	fmt.Println()
	fmt.Println("Kendaraan paling sering lewat ", commonVehicle(arr, n))

}

func catatLaluLintas(arr *Trafik, A *int) {
	var i int

	fmt.Scan(A)
	i = 1
	for i <= *A {
		fmt.Scan(&arr[i])
		i++
	}
}

func commonVehicle(arr Trafik, n int) int {
	var total1, total2, total3, total4, total5 int
	var i int
	var max urutan

	for i < n {
		if arr[i] == 1 {
			total1 += 1
			max[0] = total1
		} else if arr[i] == 2 {
			total2 += 1
			max[1] = total2
		} else if arr[i] == 3 {
			total3 += 1
			max[2] = total3
		} else if arr[i] == 4 {
			total4 += 1
			max[3] = total4
		} else if arr[i] == 5 {
			total5 += 1
			max[4] = total5
		}
		i++
	}

	return max[3]
}

func urut(arr *Trafik, n int) {
	var i, j int
	var temp int

	i = 1
	for i < n {
		j = i
		temp = arr[j]
		for j > 0 && temp > arr[j-1] {
			arr[j] = arr[j-1]
			j = j - 1
		}
		arr[j] = temp
		i = i + 1
	}
}
