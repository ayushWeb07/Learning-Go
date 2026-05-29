package main

import "fmt"

func InsertionSort(arr *[]int) {
	n := len(*arr)

	for i := 1; i < n; i++ {
		j := i

		for j > 0 && (*arr)[j-1] > (*arr)[j] {
			(*arr)[j], (*arr)[j-1] = (*arr)[j-1], (*arr)[j]
			j--
		}
	}
}

func main() {
	// create a slice
	nums := []int{5, 1, 3, 2, 4}
	fmt.Println("Before:", nums)

	InsertionSort(&nums)
	fmt.Println("After:", nums)
}
