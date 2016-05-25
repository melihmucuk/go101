package main

import (
	"errors"
	"fmt"
)

func unexportedFunc() {
	fmt.Println("This is an unexported func and you don't have to write comment on top of this function")
}

// ExportedFunc başka paketlerden çağırılabilir bir fonksiyon.
// Bu sebeple ExportedFunc için comment yazmanız gerekir.
func ExportedFunc() {
	fmt.Println("This is an exported func and you should write comment on top of this function")
}

func ping() string {
	return "pong"
}

func isEven(number int) (string, error) {
	if number%2 == 0 {
		return "Yes!", nil
	}
	return "Sorry :(", errors.New("It's an error message")
}

func f() {
	fmt.Println("hello goroutine")
}

func p(x *int) {
	*x = 10
}

func pn(x int) {
	x = 20
}
