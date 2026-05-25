package main

import (
	"fmt"
	"time"
)

type Product struct {
	id        int
	name      string
	category  string
	createdAt time.Time
}

func changeName(p *Product, newName string) {
	p.name = newName
}

func (p *Product) chageId(newId int) {
	p.id = newId
}

func main() {
	// create the struct object
	p := Product{
		id:        101,
		name:      "HyperX Quadcast",
		category:  "Mics",
		createdAt: time.Now(),
	}

	fmt.Println("id ->", p.id)
	fmt.Println("name ->", p.name)
	fmt.Println("category ->", p.category)
	fmt.Println("createdAt ->", p.createdAt)

	// copy by value
	p2 := p
	p2.category = "YT Mics"

	fmt.Println("\n", p.category, p2.category)

	// pointer with struct
	p3 := &Product{
		id:        431,
		name:      "Solocast",
		category:  "Mics",
		createdAt: time.Now(),
	}

	fmt.Println()
	fmt.Println("id ->", p3.id)
	fmt.Println("name ->", p3.name)
	fmt.Println("category ->", p3.category)
	fmt.Println("createdAt ->", p3.createdAt)

	// pass by reference
	fmt.Println()
	fmt.Println("Before:", p.name)
	changeName(&p, "Quadcast")
	fmt.Println("After:", p.name)

	// pass by reference
	fmt.Println()
	fmt.Println("Before:", p3.name)
	changeName(p3, "Lonecast")
	fmt.Println("After:", p3.name)

	// call by reference the receiver
	fmt.Println()
	fmt.Println("Before:", p.id)
	p.chageId(10101)
	fmt.Println("After:", p.id)
}
