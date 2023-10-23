package main

import "fmt"

const NMax int = 10013

type account struct {
	name    string
	accId   int
	balance float64
}

type accounts struct {
	accList [NMax]account
	total   int
}

func loadAccounts(data *accounts) {
	var i int
	var name string
	var accId int
	var balance float64

	fmt.Scan(&name, &accId, &balance)
	for name != "-1" || accId != -1 || balance != -1 {

		data.accList[i].name = name
		data.accList[i].accId = accId
		data.accList[i].balance = balance
		i++
		fmt.Scan(&name, &accId, &balance)

	}
	data.total = i
}

func sortAccounts(data *accounts) {
	var i, j int
	var temp account

	for i < data.total {
		j = i
		temp = data.accList[j]
		for j > 0 && temp.balance > data.accList[j-1].balance {
			data.accList[j] = data.accList[j-1]
			j--
		}
		data.accList[j] = temp
		i++
	}
}
func displayAccountByName(data accounts, name string) {
	var i int

	for i < data.total {
		if data.accList[i].name == name {
			fmt.Println(data.accList[i])
		} else {
			fmt.Println("NO ACCOUNT")
		}
		i++
	}
}

func main() {
	var A accounts
	var custName string
	var i int

	fmt.Println("Masukkan data : ")
	loadAccounts(&A)

	fmt.Println()
	fmt.Println("Data kamu : ")

	for i < A.total {
		fmt.Println(A.accList[i])
		i++
	}

	sortAccounts(&A)
	fmt.Println()
	fmt.Println("Data kamu setelah diurutkan : ")
	i = 0
	for i < A.total {
		fmt.Println(A.accList[i])
		i++
	}

	fmt.Println()
	fmt.Println("Customer Name ? ")
	fmt.Scan(&custName)
	for custName != "DONE" {
		displayAccountByName(A, custName)
		fmt.Println("Customer Name ? ")
		fmt.Scan(&custName)
		if custName == "DONE" {
			fmt.Println("Terima kasih sudah menggunakan aplikasi kami")
		}
	}
}
