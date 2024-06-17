package main

import "fmt"

func main() {
	var num int
	fmt.Print("Masukkan sebarang bilangan: ")
	fmt.Scan(&num)

	fmt.Println(evenOrOddNum(num))
}

func evenOrOddNum(num int) string {
	if num%2 == 0 {
		return "bilangan genap"
	} else {
		return "bilangan ganjil"
	}
}
