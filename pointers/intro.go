package main

import (
	"fmt"
	"reflect"
)

func main() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount) // Prints the address of the amount

	// Prints the type of a pionter
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))
	var myString string
	fmt.Println(reflect.TypeOf(&myString))
}