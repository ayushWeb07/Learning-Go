package main

import "fmt"

func change(ptr *int) {
	*ptr = 25
}

func main() {
	// simple & and *
	x := 10
	xAdd := &x
	fmt.Println(x, xAdd, *xAdd)

	// change value
	*xAdd = 20
	fmt.Println(x)

	// change value in function
	change(xAdd)
	fmt.Println(x)
}
