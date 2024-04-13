package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// deklarasi struct, struct person memiliki 2 properti (field), yaitu name dan age
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Retno"})

	// membuat variabel objek
	var p1 person
	p1.name = "Retno Wati"
	p1.age = 24
	fmt.Println("Name:", p1.name)
	fmt.Println("Age:", p1.age)

	// inisialisasi object struct
	var p2 = person{}
	p2.name ="Retno Wati"
	p2.age = 24
	fmt.Println("Person 2:", p2.name, "Age:", p2.age)

	var p3 =person{"Tomi Sanusi", 33}
	fmt.Println("Person 3:", p3.name, "Age:", p3.age)

	var p4 = person{name: "Afifah Nur Oktavia", age: 6}
	fmt.Println("Person 4:", p4.name, "Age:", p4.age)

	var p5 = person{name: "Latri Nur Fadhilah"}
	fmt.Println("Person 5:", p5.name, "Age:", p5.age)

	p6 := person{}
	p6.name ="Retno Wati"
	p6.age = 24
	fmt.Println("Person 6:", p6.name, "Age:", p6.age)

	p7 := person{"Tomi Sanusi", 33}
	fmt.Println("Person 7:", p7.name, "Age:", p7.age)

	p8 := person{name: "Afifah Nur Oktavia", age: 6}
	fmt.Println("Person 8:", p8.name, "Age:", p8.age)

	p9 := person{name: "Latri Nur Fadhilah"}
	fmt.Println("Person 9:", p9.name, "Age:", p9.age)
	
	p10:= person{age: 30, name: "Lina Yuniarti"}
	fmt.Println("Person 10:", p10.name, "Age:", p10.age)

	// variabel objek pointer
	s1 := student{name: "Retno", grade: 3.53}
	fmt.Println("student 1, name:", s1.name, "grade:", s1.grade)

	var s2 *student = &s1
	fmt.Println("student 2, name:", s2.name, "grade:", s2.grade)

	s2.name = "Retno Wati"
	fmt.Println("student 1, name:", s1.name, "grade:", s1.grade)
	fmt.Println("student 2, name:", s2.name, "grade:", s2.grade)

	e1 := employee{}
	e1.name = "Retno"
	e1.age = 24
	e1.address = "Gunungkidul"

	fmt.Println("name:", e1.name)
	fmt.Println("age:",e1.age)
	fmt.Println("age:", e1.person.age)
	fmt.Println("address:", e1.address)

	// pengisian nilai substruct
	person1 := person{name:"Afriska Yusuf Widyamto", age: 26}
	e2 := employee{person: person1, address: "Boyolali"}
	fmt.Println("name:", e2.name)
	fmt.Println("age:",e2.age)
	fmt.Println("address:", e2.address)

	// anonymous struct -> tidak dideklarasikan di awal sebagai tipe data baru,melainkan langsung ketika pembuatan objek
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	// deklarasi anonymous struct menggunakankeyword var
	var leader struct{
		name string
		age int
	}
	leader.name = "Afriska"
	leader.age = 27
	fmt.Println(leader.name, leader.age)

	// tipe alias
	type people = person
	p11 := people{name: "Puji Astuti", age: 56}
	fmt.Println(p11)
}

type student struct{
	name string
	grade float64
}

type employee struct{
	address string
	person // embedded struct
}
