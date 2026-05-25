package main

import (
	"fmt"
	"sync"
	"time"
)

type Food struct {
	name     string
	prepTime time.Duration
}

func prepOrder(food Food, wg *sync.WaitGroup) {
	defer wg.Done() // at the end of this function, this fires up which frees up the memory

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

	// initialize wait group
	var wg sync.WaitGroup

	for _, itm := range foods {
		// add a go routune inside the wait group
		wg.Add(1)
		go prepOrder(itm, &wg)
	}

	wg.Wait() // waits for all workers to complete
}
