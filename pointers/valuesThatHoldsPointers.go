package main

import (
	"fmt"
)

func main() {
	var myInt int
	var myIntPointer *int // Declares a variable that holds a pointer to an int
	myIntPointer = &myInt // Assign a pointer to the variable
	emeka := &myIntPointer
	fmt.Println(myIntPointer)
	fmt.Println(emeka)
}