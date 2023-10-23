package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)
	for x != y {
		(sort(x, y))
		fmt.Scan(&x, &y)
	}
}

func sort(x, y int) {
	if x < y {
		fmt.Println(x, y)
	} else {
		fmt.Println(y, x)
	}
	return
}
