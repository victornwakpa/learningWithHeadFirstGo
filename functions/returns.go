package main

import (
	"fmt"
)

// func double(number float64) float64 {
// 	return number * 2
// }
// func main() {
// 	dozen := double(6.0)
// 	fmt.Println(dozen)
// 	fmt.Println(double(5.0))
// }

// More example using the if statement
func status(grade float64) string {
	if grade < 60.0 {
		return "failing"
	}
	return "passing"
}

func main() {
	fmt.Println(status(60))
	fmt.Println(status(45))
}