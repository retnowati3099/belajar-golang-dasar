package main

import "fmt"

func main(){
	var n int
	fmt.Print("Masukkan sebarang bilangan: ")
	fmt.Scan(&n)

	fmt.Println(isPrime(n))
}

func isPrime(n int) bool{
	if n <= 1 {
		return false
	} else{
		count := 0
		for i := 1; i <= n; i++{
			if n % i == 0{
				count++
			}
		}
		return count == 2
	}
}

/*
Write the function Go that checks if a given number is a prime number. 
Output: a boolean value 'true' if the number is prime, 'false' otherwise.
*/