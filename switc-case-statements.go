package main

import "fmt"

func main() {
	point := 6
	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

	// switch case untuk banyak kondisi
	grade := 9
	switch grade {
	case 10:
		fmt.Println("Perfect")
	case 7, 8, 9:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

	// switch initializer statement
	switch score := 5; {
	case score == 10:
		fmt.Println("Perfect")
	case score >= 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
	}

	// switch case dengan gaya if else
	result := 10
	switch {
	case result == 10:
		fmt.Println("Perfect")
	case result >= 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
		fmt.Println("You need to learn more")
	}

	// keyword fallthrough dalam switch
	value := 7
	switch {
	case value == 10:
		fmt.Println("Perfect")
	case value >= 7:
		fmt.Println("Awesome")
		fallthrough
	default:
		fmt.Println("Not bad")
		fmt.Println("You need to learn more")
	}

}

/*
	Keyword fallthrough digunakan untuk memaksa proses pengecekan diteruskan ke satu case selanjutnya dengan tanpa menghiraukan nilai kondisinya, jadi satu case di pengecekan selanjutnya tersebut selalu dianggap benar meskipun aslinya salah
*/
