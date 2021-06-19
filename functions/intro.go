package main

import (
	"fmt"
)

func main() {
	sayHi("Nwakpa", 21)
}

func sayHi(victor string, age int) {
	fmt.Printf("victor %v is %d years old!", victor, age)
}

