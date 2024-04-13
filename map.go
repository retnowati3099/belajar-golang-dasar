package main

import (
	"fmt"
)

func main() {
	// membuat map menggunakan var dan :=
	var a = map[string]string{
		"anak1": "Tomi Sanusi",
		"anak2": "Retno Wati",
	}
	fmt.Printf("a\t%v\n", a)

	b := map[string]int{
		"cow":  2,
		"goat": 13,
	}
	fmt.Printf("b\t%v\n", b)

	var c map[string]string
	c = map[string]string{}
	c["fruit 1"] = "banana"
	c["fruit 2"] = "papaya"
	fmt.Printf("c\t%v\n", c)

	var countryCity = make(map[string]string)
	countryCity["Indonesia"] = "Jakarta"// inisialisasi nilai map
	countryCity["Malaysia"] = "Kuala Lumpur"
	countryCity["Filipina"] = "Manila"
	fmt.Printf("countryCity\t%v\n", countryCity)

	fruits := make(map[string]int)
	fruits["banana"] = 2
	fruits["papaya"] = 3
	fruits["apple"] = 4
	fmt.Printf("Fruits\t%v\n", fruits)

	// perbedaan deklarasi empty map menggunakan fungsi map() dan tanpa fungsi map()
	var fmake = make(map[string]string)
	var wfmake map[string]string
	
	fmt.Println(fmake == nil)
	fmt.Println(wfmake == nil)

	fmt.Println(fmake)
	fmt.Println(wfmake)

	var student = make(map[string]string)
	student["firstName"] = "Retno"
	student["lastName"] = "Wati"
	student["fullName"] = "Retno W."

	// mengakses elemen map
	fmt.Println(student["fullName"])

	// mengupdate dan menambahkan elemen map
	student["fullName"] = "Retno Wati" // update elemen
	student["origin"] = "Gunungkidul" // menambahkan elemen
	fmt.Println(student)

	// remove elemen dari map -> menggunakan fungsi delete()
	delete(student, "origin")
	fmt.Println(student)

	// mengecek elemen tertentu dari map
	v1, ok1 := student["firstName"]
	fmt.Println(v1, ok1)

	v2, ok2 := student["hobby"]
	fmt.Println(v2, ok2)

	_, ok3 :=student["fullName"]
	fmt.Println(ok3)

	// map adalah reference
	// jika dua variabel map mengacuk ke hash table yang sama, perubahan di salah satu map akan berpengaruh ke map yang lainnya
	var pet = map[string]string{
		"name": "bagong",
		"color": "brown",
		"species": "cow",
	}

	yourPet := pet
	fmt.Println(pet)
	fmt.Println(yourPet)

	yourPet["color"] = "white"
	fmt.Println(pet)
	fmt.Println(yourPet)

	// iterasi pada map
	for k, v := range pet {
		fmt.Printf("%v: %v\n", k, v)
	}
}
