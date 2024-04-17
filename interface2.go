package main

import (
	"fmt"
	"math"
)

type geometri interface {
	luas() float64
	keliling() float64
}

type persegiPanjang struct {
	panjang, lebar float64
}

type lingkaran struct {
	radius float64
}

func (p persegiPanjang) luas() float64 {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() float64 {
	return 2 * (p.panjang + p.lebar)
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.radius, 2)
}

func (l lingkaran) keliling() float64 {
	return 2 * math.Pi * l.radius
}

func main() {
	var p geometri
	p = persegiPanjang{panjang: 20, lebar: 10}
	fmt.Println("Luas persegi panjang:", p.luas())
	fmt.Println("Keliling persegi panjang:", p.keliling())

	var l geometri
	l = lingkaran{radius: 100}
	fmt.Println("Luas lingkaran:", l.luas())
	fmt.Println("Keliling lingkaran:", l.keliling())

}
