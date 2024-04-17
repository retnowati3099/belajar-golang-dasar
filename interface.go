package main

import "fmt"

type bangunDatar interface {
	luas() float64
}

type segitiga struct {
	alas   float64
	tinggi float64
}

type jajargenjang struct {
	alas   float64
	tinggi float64
}

func (s segitiga) luas() float64 {
	return (s.alas * s.tinggi) / 2
}

func (j jajargenjang) luas() float64 {
	return j.alas * j.tinggi
}

func main() {
	var bgnDatar bangunDatar
	bgnDatar = segitiga{alas: 10, tinggi: 20}
	fmt.Println("Luas segitiga:", bgnDatar.luas())

	bgnDatar = jajargenjang{alas: 10, tinggi: 20}
	fmt.Println("Luas jajargenjang:", bgnDatar.luas())
}
