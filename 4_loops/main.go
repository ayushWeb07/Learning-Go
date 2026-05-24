package main

import "fmt"

func main() {
	// simple for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	//simple while loop
	j := 1

	for j <= 5 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	//for {
	//
	//}

	fmt.Println()

	// using break to break loop flow
	for i := 1; i <= 5; i++ {
		if i == 4 {
			break
		}

		fmt.Println(i)
	}

	fmt.Println()

	// using continue to skip any iteration
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}

		fmt.Println(i)
	}

	fmt.Println()

	// using range based for loop
	for i := range 11 {
		fmt.Println(i)
	}

}
