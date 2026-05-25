package main

import "fmt"

func main() {
	// simple map creation
	winsUfc := make(map[string]int)

	winsUfc["Khabib"] = 29
	winsUfc["Conor"] = 22
	winsUfc["Islam"] = 28

	fmt.Println(winsUfc["Khabib"])
	fmt.Println(winsUfc["Conor"])
	fmt.Println(winsUfc["Islam"])
	fmt.Println()

	// check if key exist
	wins, exist := winsUfc["Mickey"]
	if exist {
		fmt.Println(wins)
	} else {
		fmt.Println("Key does not exist")
	}

	// direct initialization
	marks := map[string]int{
		"John":    55,
		"DC":      89,
		"Khamzat": 68,
	}

	fmt.Println("\n", marks, len(marks))

	// delete key
	delete(marks, "Khamzat")
	fmt.Println("\n", marks, len(marks))
	fmt.Println()

	// map iteration
	for key, val := range winsUfc {
		fmt.Println(key, "->", val)
	}
}
