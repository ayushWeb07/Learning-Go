package main

import "fmt"

func main() {
	// explicit declaration
	var name string = "John"

	var age int
	age = 25

	var isMale bool = true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMale)

	// type inference
	var work = "Entrepreneur"
	var dreamCar = 911

	fmt.Println(work)
	fmt.Println(dreamCar)

	// short syntax
	netWorth := 198.8
	isBillionaire := false
	livesIn := "Private Island"

	fmt.Println(netWorth)
	fmt.Println(isBillionaire)
	fmt.Println(livesIn)
}
