package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// mirip while
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	// keyword for tanpa argumen
	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 3 {
			break
		}
	}

	// keyword for-range
	value := "123" // string
	for _, v := range value{
		fmt.Println(v)
	}

	arr := [5]int{0, 1, 2, 3, 4} // array
	for _, v := range arr{
		fmt.Println(v)
	}

	slice := arr[0:3] // slice
	for _, v := range slice {
		fmt.Println(v)
	}

	maps := map[byte]int{
		'a': 0,
		'b': 1,
		'c': 2,
	} // map
	for _, v := range maps{
		fmt.Println(v)
	}

	// key dan value diabaikan
	for range maps{
		fmt.Println("Do")
	}

}
