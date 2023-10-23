package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	ganjil(n)
}

func ganjil(n int) {
	if n == 1 {
		fmt.Print(n, " ")
	} else {
		if n%2 == 1 {
			ganjil(n - 2)
			fmt.Print(n, " ")
		} else {
			ganjil(n - 1)
		}
	}
}
