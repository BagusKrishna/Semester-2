package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)
	for x != y {
		if x < y {
			fmt.Println(x, y)
		} else {
			fmt.Println(y, x)
		}
		fmt.Scan(&x, &y)
	}
}
