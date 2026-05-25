package main

import "fmt"

func main() {
	u := struct {
		firstName string
		lastName  string
		age       int
		isMarried bool
	}{
		"John",
		"Doe",
		29,
		false,
	}

	fmt.Println(u.firstName, u.lastName, u.age, u.isMarried)
}
