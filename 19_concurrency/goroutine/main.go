package main

import (
	"fmt"
	"time"
)

type Food struct {
	name     string
	prepTime time.Duration
}

func prepOrder(food Food) {
	fmt.Printf("Started preparing: %s...\n", food.name)
	time.Sleep(food.prepTime * time.Second)
	fmt.Printf("%s has been prepared\n\n", food.name)
}

func main() {
	foods := []Food{
		{
			name:     "Roll",
			prepTime: 5,
		},
		{
			name:     "Pasta",
			prepTime: 7,
		},
		{
			name:     "Momo",
			prepTime: 3,
		},
	}

	for _, itm := range foods {
		go prepOrder(itm)
	}

	// run infinite loop just so that the goroutines can finish their execution
	for {

	}
}
