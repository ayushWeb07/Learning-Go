package main

func main() {
	// explicit declaration
	var name string = "John"

	var age int
	age = 25

	var isMale bool = true

	println(name)
	println(age)
	println(isMale)

	// type inference
	var work = "Entrepreneur"
	var dreamCar = 911

	println(work)
	println(dreamCar)

	// short syntax
	netWorth := 198.8
	isBillionaire := false
	livesIn := "Private Island"

	println(netWorth)
	println(isBillionaire)
	println(livesIn)
}
