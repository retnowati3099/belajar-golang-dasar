package main

// menggunakan single import statement
import (
	"fmt"
	"math"
)

func main() {
	// menemukan akar pangkat dua
	fmt.Println(math.Sqrt(64))

	// menemukan akar pangkat tiga
	fmt.Println(math.Cbrt(64))

	// menemukan nilai maksimum
	fmt.Println(math.Max(64, 100))

	// menemukan nilai minimum
	fmt.Println(math.Min(64, 100))

	// menemukan sisa pembagian
	fmt.Println(math.Mod(5, 3))
}
