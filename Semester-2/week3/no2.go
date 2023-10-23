package main

import "fmt"

func main() {
	var total1, total2, total3 int
	var predicate string

	total1 = 0
	total2 = 0
	total3 = 0

	for predicate != "STOP" {
		fmt.Scan(&predicate)
		if predicate == "Sufficient" {
			total1 = total1 + 1
		} else if predicate == "Good" {
			total2 = total2 + 1
		} else if predicate == "Excellent" {
			total3 = total3 + 1
		}
	}
	fmt.Println("Total siswa dengan predikat Sufficient = ", total1)
	fmt.Println("Total siswa dengan predikat Good = ", total2)
	fmt.Println("Total siswa dengan predikat Excellent = ", total3)
}
