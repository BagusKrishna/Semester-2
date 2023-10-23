package main

import "fmt"

func main() {
	var n, fib int

	fmt.Scan(&n)
	fib = fibbo(n)
	fmt.Println(fib)
}

func fibbo(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return fibbo(n-1) + fibbo(n-2)
	}
}
