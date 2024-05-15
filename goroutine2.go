package main

import (
	"fmt"
	"time"
)

func cetakAngka(angka int) {
	for i := 0; i < 5; i++ {
		fmt.Println(angka)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go cetakAngka(1)
	go cetakAngka(2)

	time.Sleep(time.Second)
}
