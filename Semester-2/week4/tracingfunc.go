package main

import "fmt"

func main() {
	var a int = 10
	procB(&a, &a)
	fmt.Println(a)
}

func procB(b *int, c *int) {
	*c = *c + *b + *c
	*b = *b + *c - *b - *b
}
