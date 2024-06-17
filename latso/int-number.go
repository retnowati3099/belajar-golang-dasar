package main

import "fmt"

func main() {
	var num int
	fmt.Print("Masukkan sebarang bilangan: ")
	fmt.Scan(&num)

	fmt.Println(integerNumber(num))
}

func integerNumber(num int) string {
	if num > 0 {
		return fmt.Sprintf("%d adalah bilangan positif", num)
	} else if num == 0 {
		return fmt.Sprintf("%d adalah bilangan nol", num)
	} else {
		return fmt.Sprintf("%d adalah bilangan negatif", num)
	}
}
