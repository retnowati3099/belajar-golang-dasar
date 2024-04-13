package main

import "fmt"

func main(){
	var number int
	fmt.Print("Masukkan sebarang bilangan: ")
	// mengambil nilai input
	fmt.Scan(&number)

	// mencetak menggunakan fungsi Println()
	fmt.Println("Number is", number)

	fmt.Print("Menggunakan Print")
	fmt.Println("Menggunakan Println")

	var number1 int
	fmt.Print("Masukkan sebarang bilangan: ")
	fmt.Scanf("%d", &number1)
	fmt.Printf("%d", number1)
}

/*
	- Fungsi Print() untuk mencetak teks ke output screen
	- Fungsi Println() untuk mencetak teks yang akan dihasilkan dengan karaker baris baru di bagianakhir
	- Fungsi Printf() mencetak string yang diformat ke ouput screen
	- Fungsi Scan()untuk mendapatkan nilai input dari user
	- Fungsi Scanf() untuk mendapatkan nilai input menggunakan penentu format
	- Fungsi Scanln() mendapatkan nilai input hingga baris baru terdeteksi
*/