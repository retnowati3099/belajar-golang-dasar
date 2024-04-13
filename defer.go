package main

import (
	"fmt"
)

func main() {
	// defer digunakan untuk mengakhirkan eksekusi sebuah statement tepats ebelum blok fungsi selesai
	// defer fmt.Println("Hallo")
	// fmt.Println("Selamat datang")

	// statement akan tetap muncul meskipun blok kode diberhentikandi tengah jalan menggunakan return
	orderSomeFood("pizza")
	orderSomeFood("burger")

	// eksekusi defer adalah di akhir blok fungsi, bukan blok lainnya sepertu blok seleksi kondisi
	number := 3

	if number == 3{
		fmt.Println("halo 1")
		// defer fmt.Println("halo 3")
		func(){
			defer fmt.Println("halo 3")
		}()
	}
	fmt.Println("halo 2")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terima kasih, silakan tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza di tempat kami paling enak!", "\n")
		return
	}
	fmt.Println("Pesanan anda:", menu)
}
