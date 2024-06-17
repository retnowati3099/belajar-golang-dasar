package main

import "fmt"

func main() {
	var grade int
	fmt.Println("Masukkan nilai:")
	fmt.Scan(&grade)

	fmt.Println(gradeToLetterConv(grade))
}

func gradeToLetterConv(grade int) string {
	if grade >= 80 && grade <= 100{
		return fmt.Sprint("Nilai kamu A")
	} else if grade >= 70 {
		return fmt.Sprint("Nilai kamu B")
	} else if grade >= 60 {
		return fmt.Sprint("Nilai kamu C")
	} else if grade >= 50{
		return fmt.Sprint("Nilai kamu D")
	} else {
		return fmt.Sprint("Nili kamu E")
	}
}
