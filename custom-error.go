package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error){
	if strings.TrimSpace(input) == ""{
		return false, errors.New("Cannot be empty")
	}else{
		return true, nil
	}
}

func main() {
	var name string
	fmt.Print("Type your name ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid{
		fmt.Println("Hallo", name)
	} else{
		fmt.Println(err.Error())
	}
}
