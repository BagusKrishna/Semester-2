package main

import "fmt"

const MAXCALON = 74

type senator struct {
	nama     string
	nourut   int
	pengurus bool
	suara    int
}

type senatorList struct {
	daftar [MAXCALON]senator
	nCalon int
}

func main() {
	var calon senatorList
	var nama string
	var suara int

	isiTabel(&calon)
	fmt.Println("Nama calon?")
	fmt.Scan(&nama)
	suara = suaraCalon(calon, nama)
	if suara >= 0 {
		fmt.Println(nama, "mendapat", suara, "suara")
	} else {
		fmt.Println(nama, "bukan calon senator")
	}
	cetakSenatorTerpilih(calon)
}

func isiTabel(t *senatorList) {
	var i int
	fmt.Println("Masukan jumlah senator : ")
	fmt.Scan(&t.nCalon)
	fmt.Println("Masukkan daftar senator")
	i = 0
	for i < t.nCalon {
		fmt.Scan(&t.daftar[i].nama)
		fmt.Scan(&t.daftar[i].nourut)
		fmt.Scan(&t.daftar[i].pengurus)
		fmt.Scan(&t.daftar[i].suara)
		i++
	}

}

func suaraCalon(t senatorList, nama string) int {
	var i int
	var suara int = -1
	i = 1
	for i <= t.nCalon {
		if t.daftar[i].nama == nama {
			suara = t.daftar[i].suara
		}
		i++
	}
	return suara
}

func urutCalon(t *senatorList) {
	var i, j int
	var temp senator

	i = 1
	for i <= t.nCalon {
		j = i
		temp = t.daftar[j]
		for j > 0 && temp.suara > t.daftar[j-1].suara {
			t.daftar[j] = t.daftar[j-1]
			j = j - 1
		}
		t.daftar[j] = temp
		i = i + 1
	}
}

// func urutCalon(t *senatorList) {
// 	var i, j int
// 	var temp senator

// 	for i = 1; i <= t.nCalon; i++ {
// 		for j = i + 1; j <= t.nCalon; j++ {
// 			if t.daftar[i].suara > t.daftar[j].suara {
// 				temp = t.daftar[i]
// 				t.daftar[i] = t.daftar[j]
// 				t.daftar[j] = temp

// 			}
// 		}
// 	}
// }

func cetakSenatorTerpilih(t senatorList) {
	var i int = 0
	var n int = 4

	urutCalon(&t)

	// fmt.Println("urutan calon descending :")
	// for i < t.nCalon {
	// 	fmt.Println(t.daftar[i].nama, t.daftar[i].nourut, t.daftar[i].pengurus, t.daftar[i].suara)
	// 	i++
	// }

	fmt.Println("Empat calon terpilih adalah")
	for i < n {
		if t.daftar[i].pengurus != true {
			fmt.Println(t.daftar[i].nama, t.daftar[i].nourut, t.daftar[i].suara)
		} else {
			n++
		}
		i++

	}
}

// for i < t.nCalon {
// 	if t.daftar[i].nama == "Yanes" {
// 		fmt.Println(t.daftar[i])
// 	}
// 	fmt.Println(t.daftar[i].nama, t.daftar[i].nourut, t.daftar[i].pengurus, t.daftar[i].suara)
// }
