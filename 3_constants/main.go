package main

import "fmt"

func main() {
	// explicit declaration
	const name int = 18
	const age string = "Dalton"
	fmt.Println(name, age)

	// type inference
	const worksIn = "Ayush Corp."
	fmt.Println(worksIn)

	// const groups
	const (
		port          = 3000
		host          = "localhost"
		cloudProvider = "AWS"
		billAmt       = 125.89
		isTooCostly   = true
	)
	fmt.Println(port, host, cloudProvider, billAmt, isTooCostly)
}
