package main

import "fmt"

func counter() (func() int, func() int) {
	c := 0

	return func() int {
			c += 1
			return c
		}, func() int {
			c -= 1
			return c
		}
}

func main() {
	incr, decr := counter()

	fmt.Println(incr())
	fmt.Println(incr())
	fmt.Println(incr())
	fmt.Println(decr())
	fmt.Println(decr())
	fmt.Println(incr())
	fmt.Println(decr())
}
