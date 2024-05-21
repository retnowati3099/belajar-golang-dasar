package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan string)
	sayHelloTo := func(who string) {
		data := fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("Retno Wati")
	go sayHelloTo("Tomi Sanusi")
	go sayHelloTo("Afriska Yusuf Widyamto")

	messages1 := <-messages
	fmt.Println(messages1)

	messages2 := <-messages
	fmt.Println(messages2)

	messages3 := <-messages
	fmt.Println(messages3)
}
