package main

import "fmt"

func main() {
	var name = "Retno Wati"
	fmt.Printf("Tipe variabel name adalah %T\n", name)

	var fullName, age, salary = "Retno Wati", 25, 15000000.0
	fmt.Printf("fullName: %T, age: %T, salary: %T\n", fullName, age, salary)

	// deklarasi singkat
	namaLengkap, umur, gaji, softwereDeveloper := "Retno Wati", 25, 15000000.0, true

	fmt.Printf("nama lengkap: %T, umur: %T, gaji: %T, pekerjaan:%T\n", namaLengkap, umur, gaji, softwereDeveloper)
}
