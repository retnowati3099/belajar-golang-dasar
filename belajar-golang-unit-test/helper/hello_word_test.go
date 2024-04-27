package helper

import "testing"

func TestHelloWord(t *testing.T) {
	result := HelloWord("Eko")
	if result != "Hello Eko" {
		// unit test failed
		panic("Result is not Hello Eko")
	}
}

func TestHelloWordRetno(t *testing.T) {
	result := HelloWord("Retno")
	if result != "Hello Retno" {
		// unit test failed
		panic("Result is not Hello Retno")
	}
}
