package main

import (
	"fmt"
)

func main()  {
	// Structs allows you to group values with different types

	var myStruct struct {
		number 	float64	//variable with type FLoat64
		word 	string	//variable with type String
		toggle 	bool	//variable with type bool
	}
	//fmt.Printf("%#v \n", myStruct)	//Print out the struct value as it would appear in Go code

	// Assign values to struct fields
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	// Retrive and print values from struct fields
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)

	
}
