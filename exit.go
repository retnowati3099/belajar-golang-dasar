package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("hallo")
	os.Exit(1)
	fmt.Println("Selamat datang")
}

/*
	- Exit digunakan unuk menghentika program secara paksa pada saat itu juga
	- Semua statement setelah exit tidak akan dieksekusi, termask juga defer
	- Fungsi os.Exit() berada dalam package os. Fungsi ini memiliki sebuah parameter bertipe numerik yang wajib diisi. Angka yang dimasukkan akan muncul sebagai exit status ketika program berhenti
*/
