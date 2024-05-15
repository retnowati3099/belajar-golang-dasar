package main

import (
	"fmt"
	"runtime"
	"time"
)

func Print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println(i+1, message)
	}
}

func main() {
	// implementasi sederhana goroutine. Program menampilkan 8 baris teks, 4 dieksekusi dengan cara biasa, dan 4 dieksekusi sebagai goroutine baru

	runtime.GOMAXPROCS(2)

	go Print(4, "hai")
	Print(4, "hello")
}

/*
	- Goroutine sangat ringan, hanya dibutuhkan sekitar 2 kB memori saja untuk satu buah goroutine.
	- Eksekusi goroutine bersifat asynchronous, menjadikannya tidak saling tunggu dengan goroutine.
	- Eksekusi goroutine dijalankan di multi core processor, bisa ditentukan berapa banyak core yang aktif, makin banyak makin cepat.
	- Untuk menerapkan goroutine, proses yang akan dieksekusi sebagai goroutine harus dibungkus ke dalam sebuah fungsi. Pada saat pemanggilan fungsi tersebut, ditambahkan keyword go di depannya, dengan itu goroutine baru akan dibuat dengan tugas adalah menjalankan proses yang ada dalam fungsi tersebut.
	- Fungsi runtime.GOMAXPROCS(n) digunakan untuk menentukan jumlah core atau processor yang diaktifkan untuk eksekusi program
*/
