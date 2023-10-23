package main

import "fmt"

func procB(b *int, c *int) {
	*b = *b + *c
	*c = *b + *c + *b
}

func main() {
	var a int = 10
	b := &a
	c := &a
	procB(b, c)
	fmt.Println(a)
}
