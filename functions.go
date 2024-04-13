package main

import (
	"fmt"
	"math"
	"strings"
)

// fungsi tanpa paremeter dan tanpa keyword return
func SimpleFunction() {
	fmt.Println("Hello World")
}

// fungsi dengan parameter
func Add(num1 int, num2 int) {
	fmt.Println(num1 + num2)
}

// fungsi dengan return value atau nilai balik
func Sum(num1 int, num2 int) int {
	return num1 + num2
}

// fungsi dengan predefined return value
func rectangle(panjang int, lebar int) (area int) {
	area = panjang * lebar
	return
}

// fungsi multiple return
func circle(d float64) (float64, float64) {
	// hitung luas
	area := math.Pi * math.Pow(d/2, 2)

	// hitung keliling
	circumference := math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

// melewatkan address ke fungsi
func update(a *int, t *string) {
	*a = *a + 5
	*t = *t + " Wati"
	return
}

// anonymous function
var (
	Area = func(width int, height int) int {
		return width * height
	}
)

// alias skema closure
// type FilterCallback func(string)bool

// func filter(data []string, callback FilterCallback) []string{
// 	//...
// }

func main() {
	SimpleFunction()

	Add(20, 30)

	fmt.Println(Sum(100, 50))

	sum := Sum(100, 20)
	fmt.Println(sum)

	fmt.Println("Area of rectangle:", rectangle(20, 30))

	area, circumference := circle(100)
	fmt.Printf("Luas lingkaran: %.2f\n", area)
	fmt.Printf("Keliling lingkaran: %.2f\n", circumference)

	age := 24
	text := "Retno"
	fmt.Println("Before:", text, age)
	update(&age, &text)
	fmt.Println("After: ", text, age)

	fmt.Println(Area(20, 10))

	// melewatkan argumen ke dalam anonymous function
	func(width int, height int) {
		fmt.Println(width * height)
	}(20, 40)

	// fungsi yang didefinisikan untuk menerima parameter dan return value
	fmt.Printf("100 (°F) = %.2f (°C)\n", func(f float64) float64 {
		return (f - 32.0) * 5.0 / 9.0
	}(100))

	// closure function di Golang -> anonymous function yang mengakses variabel - variabel yang didefinisikan di luar body fungsi
	width := 10
	height := 100

	func() {
		area := width * height
		fmt.Println(area)
	}()

	// anonymous function mengakses variabel di setiap iterasi perulangan di dalam fungsi
	for i := 1.0; i <= 3; i++ {
		cm := func() float64 {
			return i * 100.0
		}()
		fmt.Printf("%.2f meter = %.2f centimeter\n", i, cm)
	}
	avg := calculate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(avg)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	average := calculate(numbers...)
	fmt.Println(average)

	hobbies := []string{"coding", "reading"}
	myHobbies("Retno", hobbies...)
}

// fungsi variadic -> pembuatan fungsi dengan parameter sejenis yang tidak terbatas (jumlah parameter yang disispkan ketika pemanggilan fungsi bisa berapa saja)
func calculate(num ...int) float64 {
	total := 0
	for _, v := range num {
		total += v
	}
	avg := float64(total) / float64(len(num))
	return avg
}

// fungsi dengan parameter biasa dan variadic
func myHobbies(name string, hobbies ...string) {
	hobbiesAsString := strings.Join(hobbies, ", ")

	fmt.Printf("Hello my name is %s\n", name)
	fmt.Printf("My hobbies are %s\n", hobbiesAsString)
}
