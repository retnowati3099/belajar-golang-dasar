package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return 6 * math.Pow(k.sisi, 2)
}

func (k *kubus) keliling() float64 {
	return 12 * k.sisi
}

// implementasi di fugsi main
func main() {
	var bangunRuang hitung = &kubus{5}
	fmt.Println("Luas:", bangunRuang.luas())
	fmt.Println("Keliling:", bangunRuang.keliling())
	fmt.Println("Volume", bangunRuang.volume())
}

/*
	- Interface bisa di-embed ke interface lain, dengan menuliskan nama interface yang ingin diembed ke dalam interface tujuan
*/
