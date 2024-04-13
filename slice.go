package main

import "fmt"

func main() {
	// empty slice dengan panjang 0 dan kapasitas 0
	myslice := []int{}
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	// membuat slice dari array
	var s []int = primes[1:4]
	fmt.Printf("myslice: %v\n", s)
	fmt.Printf("length: %v\n", len(s))
	fmt.Printf("capacity: %v\n", cap(s))

	// membuat slice dengan []tipedata{nilai}
	ss := []int{1, 2, 3}
	fmt.Printf("myslice: %v\n", ss)
	fmt.Printf("length: %v\n", len(ss))
	fmt.Printf("capacity: %v\n", cap(ss))

	// membuat slice dengan fungsi make()
	sss := make([]int, 3, 5)
	fmt.Printf("myslice: %v\n", sss)
	fmt.Printf("length: %v\n", len(sss))
	fmt.Printf("capacity: %v\n", cap(sss))

	// tidak mengikutkan kapasitas
	ssss := make([]int, 5)
	fmt.Printf("myslice: %v\n", ssss)
	fmt.Printf("length: %v\n", len(ssss))
	fmt.Printf("capacity: %v\n", cap(ssss))

	prices := []int{10, 20, 30}
	fmt.Printf("myslice: %v\n",prices)
	fmt.Printf("length: %v\n", len(prices))
	fmt.Printf("capacity: %v\n", cap(prices))

	// mengakses elemen dari slice
	fmt.Println(prices[0])
	fmt.Println(prices[1])

	// mengubah elemen slice
	prices[0] = 15
	fmt.Println(prices[0])
	fmt.Println(prices)

	// menambahkan elemen ke akhir slice
	prices = append(prices, 40)
	fmt.Printf("myslice: %v\n", prices)
	fmt.Printf("length: %v\n", len(prices))
	fmt.Printf("capacity: %v\n", cap(prices))

	// menambahkan sebuah slice ke slice lain
	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)
	fmt.Printf("myslice3: %v\n", myslice3)
	fmt.Printf("length: %v\n", len(myslice3))
	fmt.Printf("capacity: %v\n", cap(myslice3))

	// efisiensi memori
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("myslice3: %v\n", numbers)
	fmt.Printf("length: %v\n", len(numbers))
	fmt.Printf("capacity: %v\n", cap(numbers))

	// mengcopy angka - angka yang diperlukan
	needNum := numbers[:len(numbers)-10]
	numCopy := make([]int, len(needNum))
	copy(numCopy, needNum)
	fmt.Printf("numCopy: %v\n", numCopy)
	fmt.Printf("length: %v\n", len(numCopy))
	fmt.Printf("capacity: %v\n", cap(numCopy))

	// pengaksesan elemen slice dengan tiga indeks
	// bagian ketiga adalah kapasitas yang tidak boleh melebihi kapasitas slice yang akan dislicing
	numCopy1 := numCopy[0:4:4]
	fmt.Printf("numCopy1: %v\n", numCopy1)
	fmt.Printf("length: %v\n", len(numCopy1))
	fmt.Printf("capacity: %v\n", cap(numCopy1))
}

/*
	- Beberapa cara membuat slice di GO, yaitu menggunakan format []tipedata{nilai}, membuat slice dari array, dan menggunakan fungsi make()
*/
