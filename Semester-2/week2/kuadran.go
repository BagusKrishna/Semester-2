package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)

	if x == 0 && y == 0 {
		fmt.Println("Origin")
	} else if x >= 0 && y >= 0 {
		fmt.Println("Kuadran 1")
	} else if x < 0 && y >= 0 {
		fmt.Println("Kuadran 2")
	} else if x < 0 && y < 0 {
		fmt.Println("Kuadran 3")
	} else if x >= 0 && y < 0 {
		fmt.Println("Kuadran 4")
	}

}
