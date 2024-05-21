package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomizer1 = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer1.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	// randomizer := rand.New(rand.NewSource(10))

	// unique seed
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("Random ke-1:", randomizer.Int())
	fmt.Println("Random ke-2:", randomizer.Int())
	fmt.Println("Random ke-3:", randomizer.Int())

	// random tipe data numerikk lainnya
	fmt.Println("Random float32:", randomizer.Float32())
	fmt.Println("Random uint:", randomizer.Uint32())

	// angka random index tertentu
	fmt.Println("Random rentang 0 hingga 10: ", randomizer.Intn(10))

	// random tipe data string
	fmt.Println("Random string 4 karakter: ", randomString(4))
}

/*
	- Fungsi rand.New(rand.NewSource(10)) digunakan untuk membuat object randomizer sekaligus penentuan nilai seednya.
	- Dari onject randomizer, Int() bisa diakse, gunanya untuk men-generateangka random dalam bentuk numerik bertipe int
*/
