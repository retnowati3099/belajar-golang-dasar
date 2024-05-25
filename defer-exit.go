package main

import "fmt"

func main() {
	defer fmt.Println("Hallo")
	fmt.Println("Selamat datang")

	orderSomeDrink("coffee")
	orderSomeDrink("Juice")

	number := 2
	if number == 2{
		fmt.Println("Hallo 0")
		defer fmt.Println("Hallo 2")
	}
	fmt.Println("Hallo 1")
}

func orderSomeDrink(drink string) {
	defer fmt.Println("Terima kasih, silakan tunggu")
	if drink == "coffee" {
		fmt.Println("Coffee di tempat kami paling nikmat")
		return
	}
	fmt.Println("Pesanan Anda:", drink)
}

/*
	- Defer digunakan untuk mengakhirkan eksekusi di sebuah statement tepat sebelum blok fungsi selesai.
	- Defer bisa ditempatkan di mana saja, awal maupun akhir blok tetapi tidak mempengaruhi kapan waktu dieksekusinya, akan selalu  dieksekusi di akhir.
	- Statement defer akan tetap muncul meskipun blok kode diberhentikan di tengah jalan menggunakan return.
	- Ketika ada banyak statement yang didefer, maka seluruhnya akan dieksekusi di akhir secara berurutan.
	- Eksekusi defer adalah di akhir blok fungsi, bukan blok lainnya seperti blok seleksi kondisi
	- Exit digunakan untuk menghentikan program secara paksa, tidak seperti return yang hanya menghentikan blok kode.
*/
