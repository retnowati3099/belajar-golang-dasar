package main

import "fmt"

func main() {
	// if statement
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// if-else statement
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if-else if-else statement
	if 6 > 0 {
		fmt.Println("6 is positive")
	} else if 6 == 0{
		fmt.Println("6 is zero")
	} else {
		fmt.Println("6 is negative")
	}

	// variabel temporary pada if-else
	point := 90.0
	if result := point / 10; result == 10 {
		fmt.Println("Perfect")
	} else if result > 7.5 {
		fmt.Println("Good")
	} else {
		fmt.Println("Not Bad")
	}

	x := true
	if x{
		fmt.Println(true)
	}

	fmt.Println(Min(2, 3))
}

func Min(num1 int, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}
