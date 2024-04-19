package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func main() {
	var secret interface{}

	secret = "Retno Wati"
	fmt.Println(secret)

	secret = []string{"apple", "banana", "papaya"}
	fmt.Println(secret)

	secret = 3.14
	fmt.Println(secret)

	var data map[string]interface{}

	data = map[string]interface{}{
		"name":      "Retno Wati",
		"grade":     3.53,
		"breakfase": []string{"apple", "manggo", "banana"},
	}
	fmt.Println(data)

	var data2 map[string]any

	data2 = map[string]any{
		"name":      "Retno Wati",
		"grade":     3.53,
		"breakfase": []string{"apple", "manggo", "banana"},
	}
	fmt.Println(data2)

	// casting variabel interface kosong
	var something any

	something = 2
	var num = something.(int) * 10
	fmt.Println(something, "multiplied by 10 is :", num)

	something = []string{"apple", "banana", "papaya"}
	var fruits = strings.Join(something.([]string), ", ")
	fmt.Println(fruits, "are my favorite fruits")

	// casting variabel interface kosong ke objek pointer
	var human any = &person{name: "Retno", age: 24}
	name := human.(*person).name
	fmt.Println(name)

	// kombinasi slice, map, dan interface
	var buah = []interface{}{
		map[string]interface{}{"name": "banana", "total": 10},
		[]string{"mango", "banana", "papaya"},
		"orange",
	}

	for _, each := range buah{
		fmt.Println(each)
	}
}

/*
	- Interface kosong atau empty interface dinotasikan dengan interface{} atau any
	- Variabel bertipe interface kosong bisa menampung segala jenis data, bahkan array, pointer, apapun.
	- Tipe data dengan konsep ini biasa disebut dengan dynamic typing
	- Tipe any adalah alias dari interface{}, keduanya adalah sama
*/
