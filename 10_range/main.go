package main

import "fmt"

func main() {
	// range on arrays
	nums := [5]int{31, 15, 59, 21, 42}

	fmt.Println("Array:-")
	for idx, val := range nums {
		fmt.Println(idx, "->", val)
	}

	// range on map
	marks := map[string]int{
		"John":    55,
		"DC":      89,
		"Khamzat": 68,
	}

	fmt.Println("\nMap:-")
	for key, val := range marks {
		fmt.Println(key, "->", val)
	}

	// range on string
	fmt.Println("\nString:-")
	for idx, char := range "Porsche 911" {
		fmt.Println(idx, "->", string(char))
	}
}
