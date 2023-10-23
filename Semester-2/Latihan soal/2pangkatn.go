package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	fmt.Println(pangkat(n))
}

func pangkat(a int) int {
	if a == 0 {
		return 1
	} else {
		return 2 * pangkat(a-1)
	}
}
