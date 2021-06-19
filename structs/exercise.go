package main

import (
	"fmt"
)

func main() {
	var subscriber struct {
		Name	string
		Rate	float64
		Active	bool
	}
	subscriber.Name 	= "Aman Singh"
	subscriber.Rate 	= 4.99
	subscriber.Active	= true

	fmt.Println("Name: ", subscriber.Name)
	fmt.Println("Monthly rate: ", subscriber.Rate)
	fmt.Println("Active? ", subscriber.Active)
}