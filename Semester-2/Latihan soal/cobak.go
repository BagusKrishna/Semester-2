package main

import "fmt"

// func main() {
// 	var fah, cel float64
// 	var n int

// 	fmt.Scan(&n)
// 	for i := 1; i <= n; i++ {
// 		fmt.Scan(&cel)
// 		fah = farenwoi(cel)
// 		fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", cel, fah)
// 	}
// }

// func farenwoi(c float64) float64 {
// 	return (9*c)/5 + 32
// }
func main() {
	var a int = 5
	var tot int
	tot = 1
	for i := 1; i <= a; i++ {
		i = i * i
	}
	fmt.Println(tot)
}
