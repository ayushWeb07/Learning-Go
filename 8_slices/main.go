package main

import "fmt"

func main() {
	// simple slice
	nums := []int{1, 5, 2, 4, 3}
	fmt.Println(nums, len(nums), cap(nums))

	// using make()
	arr := make([]int, 3, 5)
	fmt.Println(arr, len(arr), cap(arr))

	// append the values
	arr = append(arr, 5, 7)
	fmt.Println(arr, len(arr), cap(arr))

	// append more values to exceed the cap.
	arr = append(arr, 3, 1)
	fmt.Println(arr, len(arr), cap(arr))

	// empty slice
	ns := []int{}
	fmt.Println(ns, len(ns), cap(ns))
	ns = append(ns, 1, 2, 3)
	fmt.Println(ns, len(ns), cap(ns))

	// copy slices
	n1 := make([]int, 0, 5)
	n1 = append(n1, 4, 3, 1)

	n2 := make([]int, len(n1), cap(n1))
	copy(n2, n1)

	fmt.Println("n1", n1, len(n1), cap(n1))
	fmt.Println("n2", n2, len(n2), cap(n2))
	fmt.Println()

	// slice operations
	n := []int{5, 1, 3, 2, 4}
	fmt.Println(n)
	fmt.Println(n[:])
	fmt.Println(n[:3])
	fmt.Println(n[2:])
	fmt.Println(n[1:4])
}
