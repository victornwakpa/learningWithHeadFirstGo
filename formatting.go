// %f | Floating-point number
// %d | Decimal interger
// %s | String
// %t | Boolean
// %v | Any value (chooses an appropriate format basedon the supplied value's type)
// %#v| Any value, formatted as it would appear in Go program code
// %T Type of the supplied value (int, string, etc.)
// %% | A literal percent sign

package main

import (
	"fmt"
)

func main() {
	// The first argument or parameter in a Printf function must be a string
	fmt.Printf("A float %f \n", 3.1415)
	fmt.Printf("An integer: %d \n", 15)
	fmt.Printf("A string: %s \n", "Hello")
	fmt.Printf("A boolean: %t \n", false)
	fmt.Printf("Values: %v %v %v \n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v \n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T \n", 1.2, "hello", true)
	fmt.Printf("Percent sign: %% \n")
}