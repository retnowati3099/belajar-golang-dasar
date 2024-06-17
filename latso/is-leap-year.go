package main

import "fmt"

func main() {
	var year int
	fmt.Print("Masukkan sebarang tahun: ")
	fmt.Scan(&year)

	fmt.Println(isLeapYear(year))
}

func isLeapYear(year int) string {
	if year % 4 == 0{
		if year % 100 == 0{
			if year % 400 == 0 {
				return fmt.Sprintf("%d tahun kabisat", year)
			} else{
				return fmt.Sprintf("%d bukan tahun kabisat", year)
			}
		} else{
			return fmt.Sprintf("%d tahun kabisat", year)
		}
	} else{
		return fmt.Sprintf("%d bukan tahun kabisat", year)
	}
}
