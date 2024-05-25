package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hallo")
	os.Exit(1)
	fmt.Println("Selamat datang")
}

/*
	- Exit digunakan untukmenghentikan programsecara paksa pada saat itu juga.
	- Semua statement setelah exittidak akan dieksekusi, termasuk defer.
	- Fungsi os.exit() berada dalam package os dan memiliki sebuah parameter bertipe numerik yang wajib diisi. Angka yang dimasukkan akan muncul sebagai exit status keika program berhenti.
*/
