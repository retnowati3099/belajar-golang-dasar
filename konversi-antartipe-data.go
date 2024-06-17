package main

import (
	"fmt"
	"strconv"
)

func main() {
	// konversi menggunakan strconv
	// fungsi strconv.Atoi()
	var str = "1999"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num)
	}
	var str2 = "perempuan"
	var num2, err2 = strconv.Atoi(str2)
	// fmt.Println(num2,"&", err2)
	if err2 == nil {
		fmt.Println(num2)
	} else{
		fmt.Println(err2)
	}

	// fungsi strconv.Itoa()
	var num3 = 30
	var str3 = strconv.Itoa(num3)
	fmt.Println(str3)

	// fungsi strconv.ParseInt()
	// string "1999" dikonversi ke basis 10 dengan tipe data hasil adalah int8
	str4 := "1999"
	num4, err4 := strconv.ParseInt(str4, 10, 8)
	if err4 == nil {
		fmt.Println(num4)
	} else {
		fmt.Println(err4)
	}

	// string "1111" dikonversi ke basis 2 (biner) dengan tipe data hasil adalah int8
	str5 := "1111"
	num5, err5 := strconv.ParseInt(str5, 2, 8)
	if err5 == nil {
		fmt.Println(num5)
	} else{
		fmt.Println(err5)
	}

	// string "B" dikonversi ke basis 16 (heksadesimal) dengan tipe data hasil konversi adalah int8
	str6 := "B"
	num6, err6 := strconv.ParseInt(str6, 16, 8)
	if err6 == nil {
		fmt.Println(num6)
	} else {
		fmt.Println(err6)
	}
}

/*
	- Fungsi strconv.atoi() digunakan untuk konversi data dari tipe string ke int
	-Fungsi strconv.atoi()menghasilkan 2 buah nilai kembalian, yaitu hasil konversi dan error (jika konversi sukses, maka error bernilai nil)
	- Fungsi strconv.Itoa() merupakan kebalikan dari strconv.Atoi(), berguna untuk konversi int ke string
	- Fungsi strconv.ParseInt() digunakan untuk konversi string berbentuk numerik dengan basis tertentu ke tipe numerik non-desimal dengan lebar data bisa ditentukan
*/
