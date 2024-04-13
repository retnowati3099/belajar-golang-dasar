package main

import (
	"fmt"
	"strconv"
)

func main() {
	// konversi di antara string dan integer
	// konversi string ke int
	var s string = "64"
	v, _ := strconv.Atoi(s)
	fmt.Println(v)

	// konversi int ke string
	var i int = 64
	str := strconv.Itoa(i)
	fmt.Println(str)

	// konversi di antara int dan float
	// konversi float ke int -> kehilangan presisi
	ff := 3.14
	fmt.Println(int(ff))

	// konversi int ke float
	ii := 3
	fmt.Println(float64(ii))

	// konversi float ke uint
	fmt.Println(uint(ff))

	// konversi string dan byte
	// konversi string ke byte
	ss := "Hello World"
	by := []byte(ss)
	fmt.Println(by)

	// konversi byte ke string
	fmt.Println(string(by))

	// konversi type saat pembagian
	res1 := 6/3 // keduanya int, maka res1 int
	res2 := 6.3 / 3 // float dan int, maka res2 float
	fmt.Println(res1, res2)
}
