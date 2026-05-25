package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	val int
	mu  sync.Mutex
}

func (c *Counter) incr(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		c.mu.Unlock() // release the lock
	}()

	c.mu.Lock() // lock the resource
	c.val += 1  // modifying the resource
}

func main() {
	c := Counter{val: 0}
	wg := sync.WaitGroup{}

	fmt.Println("c:", c.val)

	// increment 50 times
	for range 50 {
		wg.Add(1)
		go c.incr(&wg)
	}

	wg.Wait()

	// now print value
	fmt.Println("c:", c.val)
}
