package main

import (
	"fmt"
)

// func paintNeeded(width float64, height float64) {
// 	area := width * height
// 	fmt.Printf("%.2f liters needed \n", area/10.0)
// }

// func main() {
// 	paintNeeded(4.2, 3.0)
// 	paintNeeded(5, 8)
// 	// functionA(4, 5)
// 	// functionB(5, 4)
// }

// More examples
// func functionA(a int, b int) {
// 	fmt.Println(a + b)
// }
// func functionB(a int, b int) {
// 	fmt.Println(a * b)
// }

// Using returns in our paint example
func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area/10.0
}

func main() {
	var amount, total float64
	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("%.2f litres needed \n", amount)
	total += amount
	amount = paintNeeded(5, 8)
	fmt.Printf("%.2f litres needed \n", amount)
	total += amount
	fmt.Println("total: ", total)
}
