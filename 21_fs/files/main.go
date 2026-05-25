package main

import (
	"fmt"
	"os"
)

func main() {
	// open file
	f, err := os.Open("21_fs\\files\\logs.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// close file
	defer f.Close()

	// print the file's stat
	info, err := f.Stat()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", info.Name())
	fmt.Println("IsDir:", info.IsDir())
	fmt.Println("Size:", info.Size())

	// read file content
	data, err := os.ReadFile("21_fs\\files\\logs.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nData:-\n", string(data))

	// creating files
	newFile, err := os.Create("21_fs\\files\\go.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer newFile.Close()

	newFile.WriteString("Hi Go, you're awesome!")
}
