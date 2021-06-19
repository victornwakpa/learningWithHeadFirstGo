package main

import (
	"fmt"
	// "time"
	"strings"
)

func main() {
	broken := "G# R#CKS!"
	replacer := strings.NewReplacer("#", "O") //this returns a strings.Replacer that's set up to replace every "#" with "O"
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}

// func main() {
// 	var date time.Time = time.Now()
// 	year, month, day := date.Year(), date.Month(), date.Day()
// 	fmt.Println(year, month, day)
// 	fmt.Println(date)
// }

