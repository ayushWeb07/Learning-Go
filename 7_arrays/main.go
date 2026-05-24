package main

import "fmt"

func main() {
	// simple array creation
	var nums [5]int
	size := len(nums)

	nums[0] = 2
	nums[1] = 5
	nums[2] = 1
	nums[3] = 3
	nums[4] = 4

	fmt.Println("Size: ", size)

	fmt.Println("\nFirst Array:-")

	for i := 0; i < size; i++ {
		fmt.Println(nums[i])
	}

	// array initialization
	arr := [5]int{12, 15, 8, 18, 9}

	fmt.Println("\nSecond Array:-")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("\nFirst Array once again:-")

	// range based
	for i := range size {
		fmt.Println(nums[i])
	}
}
