package main

import "fmt"

func main() {
	// boolean
	var boolean bool = false
	var boolean1 = false
	var boolean2 bool
	boolean3 := false
	fmt.Println(boolean, boolean1, boolean2, boolean3)

	// integer
	var integer int = 5
	var integer8 int8 = 127 
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647 // rune alias untuk int32
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer, integer8, integer16, integer32, integer64)

	var uinteger uint = 4294967295
	var uinteger8 uint8 = 255 // byte alias untuk int8
	var uinteger16 uint16 = 65535
	var uinteger32 uint32 = 4294967295
	var uinteger64 uint64 = 18446744073709551615
	fmt.Println(uinteger, uinteger8, uinteger16, uinteger32, uinteger64)

	// float
	// float32
	var float1_32 float32 = 123.75
	var float2_32 float32 = 3.4e5
	fmt.Println(float1_32, float2_32)

	// float64
	var float1_64 float32 = 123.75
	var float2_64 float32 = 3.4e5
	fmt.Println(float1_64, float2_64)

	// string
	var str1 string = "Hello"
	var str2 = "Hello"
	var str3 string
	str4 := "Hello"
	fmt.Println(str1, str2, str3, str4)

	fmt.Printf("Type: %T, Value: %v\n", str1, str1)
}

/*
	- Tipe data di Golang dikategorikan menjadi 2, yaitu basic types dan composite types.
*/