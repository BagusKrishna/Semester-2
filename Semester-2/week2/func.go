package main

import "fmt"

func segitiga(a, b, c int) string {
	var bentuk string
	if a+b > c && b+c > a && a+c > b {
		bentuk = "Segitiga"
	} else {
		bentuk = " Bukan Segitiga"
	}
	return bentuk
}

func main() {
	var bentuk string
	var a, b, c int

	fmt.Scan(&a, &b, &c)
	bentuk = segitiga(a, b, c)
	fmt.Println(bentuk)
}
