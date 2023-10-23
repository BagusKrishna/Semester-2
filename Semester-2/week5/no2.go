package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	jumlahBus(n, m)

}

func jumlahBus(penumpang, kapasitas int) int {
	var bus int
	if penumpang%kapasitas == 0 {
		bus = penumpang / kapasitas
	} else {
		bus = penumpang/kapasitas + 1
	}
	fmt.Println(bus, "bus")
	return bus
}
