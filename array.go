package main

import (
	"fmt"
)

func main() {
	// membuat array dan inisialisasi nilai array
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"

	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	// inisialisasi cara horisontal
	words := [2]string{"Hello", "World"}
	fmt.Println(words[0], words[1])
	fmt.Println(words)

	// inisialisasi cara vertikal
	sayHello := [2]string{
		"Hello",
		"World",
	}
	fmt.Println(sayHello[0], sayHello[1])
	fmt.Println(sayHello)

	// inisialisasi nilai awal array tanpa jumlah elemen
	var numbers = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Data array:", numbers)
	fmt.Println("Panjang array:", len(numbers))

	// alokasi elemen array menggunakan keyword make
	var animals = make([]string, 2)
	animals[0] = "cow"
	animals[1] = "goat"
	fmt.Println(animals[0], animals[1])
	fmt.Println(animals)

	// array multidimensi
	// ditulikan jumlah datanya
	var numbers1 = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	fmt.Println(numbers1)

	// tidak dituliskan julah datanya
	var numbers2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(numbers2)

	// perulangan elemen array
	var fruits = [4]string{"apple", "grape", "banana", "papaya"}

	// menggunakan keyword for
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Elemen %d : %s\n", i, fruits[i])
	}

	// menggunakan keyword for-range
	for i, v := range fruits {
		fmt.Printf("Elemen %d : %s\n", i, v)
	}

	// penggunaan variabel underscore (pengangguran)
	for _, v := range fruits {
		fmt.Printf("%s\n", v)
	}

	// jika yang dibutuhkan hanya indeks
	for i, _ := range fruits {
		fmt.Println(i)
	}

	for i := range fruits {
		fmt.Println(i)
	}
}
