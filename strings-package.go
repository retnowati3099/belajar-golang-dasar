package main

import (
	"fmt"
	"strings"
)

func main() {
	// konversi string ke lowercase
	lower := strings.ToLower("GOLANG STRINGS")
	fmt.Println(lower)

	// konversi string ke uppercase
	upper := strings.ToUpper("golang strings")
	fmt.Println(upper)

	// membuat slice array
	stringSlice := []string{"I love", "Go Programming"}

	// join elemen slice dengan spasi 
	joinedString := strings.Join(stringSlice, " ")
	fmt.Println(joinedString)

	// melakukan pengecekan apakah dua string sama
	compare1 := strings.Compare("matematika", "informatika")
	fmt.Println(compare1)

	compare2 := strings.Compare("fisika", "informatika")
	fmt.Println(compare2)

	compare3 := strings.Compare("informatika", "informatika")
	fmt.Println(compare3)

	// melakukan pengecekan jika suatu string berisi substring
	contain1 := strings.Contains("Retno Wati", "Wati")
	fmt.Println(contain1)

	contain2 := strings.Contains("Retno Wati", "Wait")
	fmt.Println(contain2)

	// menghitung jumlah karakter tertentu(parameter kedua) dari sebuah string (parameter pertama). casesensitive
	count := strings.Count("informatika", "i")
	fmt.Println(count)

	// mendeteksi apakah string (parameter pertama) diawali string tertentu(parameter kedua)
	isPrefix1 := strings.HasPrefix("Retno Wati", "Re")
	fmt.Println(isPrefix1)

	isPrefix2 := strings.HasPrefix("Retno Wati", "Wa")
	fmt.Println(isPrefix2)

	// mendeteksi apakah sebuah string (parameter pertama) diakhiri string tertentu (parameter kedua)
	isSuffix1 := strings.HasSuffix("Retno Wati", "ati")
	fmt.Println(isSuffix1)

	isSuffix2 := strings.HasSuffix("Retno Wati", "no")
	fmt.Println(isSuffix2)

	// mengganti bagian dari string dengan string tertentu
	text := "banana"
	find := "a"
	replaceWith := "o"

	newText1 := strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1)

	newText2 := strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2)

	newText3 := strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3)

	// mengulang string (parameter pertama) sebanyak data yang ditentukan (parameter kedua)
	str := strings.Repeat("la", 3)
	fmt.Println(str)

	// memisah string (parameter pertama) dengan tanda pemisah bisa ditentukan sendiri (parameter kedua). Hasilnya berupa slice string
	string1 := strings.Split("A strong women", " ")
	fmt.Println(string1)

	string2 := strings.Split("Retno", "")
	fmt.Println(string2)

	// mencari posisi indeks sebuah string (parameter kedua) dalam string (parameter pertama). Jika ditemukan dua substring, maka yang diambil adalah yang pertama
	index1 := strings.Index("Retno Wati", "no")
	fmt.Println(index1)

	index2 := strings.Index("Retno Wati", "b")
	fmt.Println(index2)
}
