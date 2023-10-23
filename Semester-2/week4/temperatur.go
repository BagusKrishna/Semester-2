package main

import "fmt"

func main() {
	var C float64

	fmt.Scan(&C)
	fmt.Println(C)
	R(C)
	F(C)
	K(C)

}
func R(c float64) {
	var r float64
	r = c * (4 / 5)
	fmt.Println("Reamur :", r)

}

func F(c float64) {
	var f float64
	f = c*(9/5) + 32
	fmt.Println("Farenheit:", f)
}

func K(c float64) {
	var k float64
	k = c + 273.15
	fmt.Println("Kelvin:", k)
}
