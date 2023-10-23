package main

import "fmt"

func main() {
	var a, n int
	a = 1

	fmt.Scan(&n)
	pola(a, n)
	fmt.Println("KAMU KONTOL !!!")
	fmt.Println("      O ")
	fmt.Println("    / | \\")
	fmt.Println("      | ")
	fmt.Println("  €==3| ")
	fmt.Println("     / \\")
}

func pola(i, n int) {
	var j int = 1
	if i == n {
		for i := 1; i <= n; i++ {
			fmt.Print("*")
		}

	} else {
		for j <= i {
			fmt.Print("*")
			j++
		}
		fmt.Println()
		pola(i+1, n)
	}
}
