package main

import "fmt"

// deklarasi banyak variabel dengan keyword var di luar fungsi
var k bool

var l bool = true
var m = false

var (
	d int
	e int = 5
	f string = "huruf"
	g bool = true
	h float64 = 2.2
)

func main() {
	// deklarasi variabel dengan keyword var di dalam fungsi
	var i string
	var j int
	var n float64
	fmt.Println(i, j, k, n)

	fmt.Println(l, m)

	// inisialisasi variabel
	var student1 string = "Retno"
	fmt.Println(student1)

	var student2 = "Afifah"
	fmt.Println(student2)

	var student3 string
	student3 = "Latri"
	fmt.Println(student3)

	// deklarasi variable dengan tanda :=
	fruit := "apple"
	fmt.Println(fruit)

	// deklarasi multi variabel
	var one, two, three string
	one, two, three = "satu", "dua", "tiga"
	fmt.Println(one, two, three)

	var four, five, six string = "empat", "lima", "enam"
	fmt.Println(four, five, six)

	var seven, eight, nine = "tujuh", "delapan", "sembilan"
	fmt.Println(seven, eight, nine)

	var first, second = "pertama", 2
	fmt.Println(first, second)

	ten, eleven, twelve := "sepuluh", "sebelas", "duabelas"
	fmt.Println(ten, eleven, twelve)

	num1, isFriday, twoPointwo, say := 1, false, 2.2, "hello"
	fmt.Println(num1, isFriday, twoPointwo, say)

	var(
		a int
		b int = 1
		c  string = "hello"
	)
	fmt.Println(a, b, c)

	fmt.Println(d, e, f, g, h)

	// variabel underscore (_), nilai tidak akan digunakan
	_ = "belajar Golang itu mudah lho"

	// deklarasi variabel dengan keyword new untuk membuat variabel pointer dengan tipe data tertentu
	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)
}


/*
	- Keyword var dapat digunakan di dalam dan di luar fungsi
	- Dengan keyword var, deklarasi variabel dan pemberian nilai dapat dilakukan secara terpisah
	- := hanya dapat digunakan di dalam fungsi
	- Dengan tanda := delarasi variabel dan pemberian nilai tidak dapat dilakukan secara terpisah (harus dilakukan di baris yang sama)
*/