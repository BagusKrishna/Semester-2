package main

import "fmt"

func main() {
	var u int
	fmt.Scan(&u)
	duid(&u)
}

func duid(u *int) {
	var lp, ld, ls int
	if *u >= 10000 {
		lp = *u / 10000
		*u = *u % 10000
		fmt.Println("sisa = ", *u)
	}
	if *u >= 2000 {
		ld = *u / 2000
		*u = *u % 2000
		fmt.Println("sisa = ", *u)
	}
	if *u >= 1000 {
		ls = *u / 1000
		*u = *u % 1000
	}
	fmt.Println(lp, "lembar pecahan Rp.10.000")
	fmt.Println(ld, "lembar pecahan Rp.2.000")
	fmt.Println(ls, "lembar pecahan Rp.1.000")
}
