package main

func main() {
	// explicit declaration
	const name int = 18
	const age string = "Dalton"
	println(name, age)

	// type inference
	const worksIn = "Ayush Corp."
	println(worksIn)

	// const groups
	const (
		port          = 3000
		host          = "localhost"
		cloudProvider = "AWS"
		billAmt       = 125.89
		isTooCostly   = true
	)
	println(port, host, cloudProvider, billAmt, isTooCostly)
}
