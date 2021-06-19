package main

import (
	"fmt"
	// "errors"
	// "log"
)

// func main() {
// 	// err := errors.New("height can't be negative")
// 	err := fmt.Errorf("a height of %.2f is invalid", -2.33333)
// 	fmt.Println(err)
// 	fmt.Println(err.Error())
// 	// log.Fatal(err)
// }

func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

func main() {
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)
}