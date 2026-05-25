package main

import (
	"fmt"
	"time"
)

type Food struct {
	name     string
	prepTime int
}

func prepOrder(food Food, foodChan chan<- string) {
	fmt.Printf("Started preparing: %s...\n", food.name)
	time.Sleep(time.Duration(food.prepTime) * time.Second)
	fmt.Printf("\n%s has been prepared", food.name)

	// send food name to the channel
	foodChan <- food.name
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

	// create a buffered channel that contains the food names
	foodChan := make(chan string, len(foods))

	for _, itm := range foods {
		go prepOrder(itm, foodChan)
	}

	// receive food names from the channel
	for range foods {
		name := <-foodChan
		fmt.Printf("\n ~Name: %s~ \n", name)
	}

}
