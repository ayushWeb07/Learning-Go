package main

import (
	"fmt"
	"sync"
	"time"
)

type Food struct {
	name     string
	prepTime int
}

func prepOrder(food Food, foodChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

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

	// create a wait group to ensure all data is sent to channel
	wg := sync.WaitGroup{}

	for _, itm := range foods {
		wg.Add(1)
		go prepOrder(itm, foodChan, &wg)
	}

	// close channel once all data is sent to channel
	go func() {
		defer close(foodChan)
		wg.Wait()
	}()

	// receive food names from the channel
	for name := range foodChan {
		fmt.Printf("\n ~Name: %s~ \n", name)
	}

}
