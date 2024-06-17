package main

import "fmt"

func main() {
	var num int
	fmt.Print("Masukkan bilangan (1-7): ")
	fmt.Scan(&num)

	fmt.Println(days(num))
}

func days(num int) string{
	switch num{
	case 1:
		return fmt.Sprint("Senin")
	case 2:
		return fmt.Sprint("Selesai")
	case 3:
		return fmt.Sprint("Rabu")
	case 4:
		return fmt.Sprint("Kamis")
	case 5:
		return fmt.Sprint("Jumat")
	case 6:
		return fmt.Sprint("Sabtu")
	case 7:
		return fmt.Sprint("Minggu")
	default:
		return fmt.Sprint("Bukan input hari")
	}
}