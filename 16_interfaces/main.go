package main

import "fmt"

type Card interface {
	deposit(amt float32)
	withdraw(amt float32)
	printDetails()
}

type CreditCard struct {
	name    string
	balance float32
}

func (c *CreditCard) deposit(amt float32) {
	c.balance += amt
	fmt.Println("Depositing into credit card...")
}

func (c *CreditCard) withdraw(amt float32) {
	if c.balance >= amt {
		c.balance -= amt
		fmt.Println("Withdrawing from credit card...")
	} else {
		fmt.Println("Insufficient funds!")
	}
}

func (c *CreditCard) printDetails() {
	fmt.Println("Name:", c.name, "| Amt ($):", c.balance, "\n")
}

type DebitCard struct {
	name    string
	balance float32
}

func (d *DebitCard) deposit(amt float32) {
	d.balance += amt
	fmt.Println("Depositing into debit card...")
}

func (d *DebitCard) withdraw(amt float32) {
	if d.balance >= amt {
		d.balance -= amt
		fmt.Println("Withdrawing from debit card...")
	} else {
		fmt.Println("Insufficient funds!")
	}
}

func (d *DebitCard) printDetails() {
	fmt.Println("Name:", d.name, "| Amt ($):", d.balance, "\n")
}

func main() {
	// create credit card object
	c := CreditCard{
		name:    "Visa Credit Card",
		balance: 100,
	}

	c.printDetails()

	c.deposit(50)
	c.printDetails()

	c.withdraw(20)
	c.printDetails()

	c.withdraw(200)
	c.printDetails()

	// create debit card object
	d := DebitCard{
		name:    "BGVB Debit Card",
		balance: 80,
	}

	d.printDetails()

	d.deposit(20)
	d.printDetails()

	d.withdraw(50)
	d.printDetails()

	d.withdraw(20)
	d.printDetails()
}
