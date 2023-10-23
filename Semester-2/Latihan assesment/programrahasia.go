package main

import "fmt"

const B int = 6
const K int = 13

type typeOne [K]int
type typeTwo [B][B]rune

func inputData(data *typeOne) {
	var x int
	x = 1
	for x <= K {
		fmt.Scan(&data[x])
		x++
	}

}

func inputMat(mat *typeTwo) {
	var i, j int
	i = 1
	for i <= B {
		j = 1
		for j <= B {
			fmt.Scan(&mat[i][j])
			j++
		}
		i++
	}
}

func outputMat(datA, datB *typeOne, mat typeTwo) {
	var u int
	u = 1
	for u <= K {
		fmt.Println(mat[datA[u]][datB[u]])
		u++
	}
}

func main() {
	var tabA, tabB typeOne
	var tabM typeTwo

	inputData(&tabA)
	inputData(&tabB)
	inputMat(&tabM)
	outputMat(&tabA, &tabB, tabM)

}
