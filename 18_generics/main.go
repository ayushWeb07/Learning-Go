package main

import "fmt"

func printSlice[T any](arr []T) {
	for _, val := range arr {
		fmt.Println(val)
	}
}

type Number interface {
	int | float64
}

func add[T Number](a, b T) T {
	return a + b
}

type Rect[T Number] struct {
	length  T
	breadth T
}

func main() {
	// pass all kinds of slice to printSlice()
	iSl := []int{5, 1, 4, 2, 3}
	printSlice(iSl)
	fmt.Println()

	fSl := []float32{12.5, 10.6, 42.9, 24.1, 32.3}
	printSlice(fSl)
	fmt.Println()

	bSl := []bool{true, false, true}
	printSlice(bSl)
	fmt.Println()

	// pass int and float to add()
	fmt.Println("7 + 25 =", add(7, 25))
	fmt.Println("12.4 + 2.5 =", add(12.4, 2.5))

	iRect := Rect[int]{
		length:  18,
		breadth: 20,
	}

	fRect := Rect[float64]{
		length:  1.5,
		breadth: 2.1,
	}

	fmt.Println("\n", iRect, fRect)

}
