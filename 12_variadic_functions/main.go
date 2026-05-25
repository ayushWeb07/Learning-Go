package main

import "fmt"

func sum(nums ...int) int {
	tot := 0

	for _, n := range nums {
		tot += n
	}

	return tot
}

func main() {

	// pass numbers spreaded
	tot1 := sum(1, 2, 5, 3, 4)
	fmt.Println(tot1)

	// pass slice
	nums := []int{10, 20, 30}
	tot2 := sum(nums...)
	fmt.Println(tot2)
}
