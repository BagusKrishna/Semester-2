package main

import "fmt"

func main() {
	var ipk float64
	var masaStudi int
	var publikasi bool
	fmt.Scan(&ipk, &masaStudi, &publikasi)

	if cumlaude(ipk, masaStudi, publikasi) {
		fmt.Println("Cumlaude")

	}
	if sangatMemuaskan(ipk, masaStudi, publikasi) {
		fmt.Println("Sangat Memuaskan")

	}
	if memuaskan(ipk, masaStudi, publikasi) {
		fmt.Println("Memuaskan")

	}
}

func cumlaude(ipk float64, masaStudi int, publikasi bool) bool {
	return 3.51 <= ipk && ipk <= 4.00 && masaStudi <= 8 && publikasi
}

func sangatMemuaskan(ipk float64, masaStudi int, publikasi bool) bool {
	return 2.71 <= ipk && ipk <= 3.5 && masaStudi <= 14
}

func memuaskan(ipk float64, masaStudi int, publikasi bool) bool {
	return 2.00 <= ipk && ipk <= 2.75 && masaStudi <= 14
}
