package main

import "fmt"

func main() {
	// simple, if else-if ladder
	n := 0

	if n > 0 {
		fmt.Println("+ve number")
	} else if n < 0 {
		fmt.Println("-ve number")
	} else {
		fmt.Println("0")
	}

	fmt.Println()

	// nested if else ladder
	haveDebitCard := true
	cardBalance := 100
	amtRequested := 150

	if !haveDebitCard {
		fmt.Println("Get a debit card first!")
	} else {
		if cardBalance >= amtRequested {
			fmt.Println("Transaction successfull!")
		} else {
			fmt.Println("Insufficient funds")
		}
	}

	fmt.Println()

	// && operator
	age := 18
	hasLiscense := true
	isDrunk := false

	if age >= 18 && hasLiscense && !isDrunk {
		fmt.Println("You're fit to drive!")
	} else {
		fmt.Println("Not eligible to drive!")
	}

	fmt.Println()

	// || operator
	day := "Sun"

	if day == "Sun" || day == "Sat" {
		fmt.Println("Weekend!")
	} else {
		fmt.Println("Weekday!")
	}
}
