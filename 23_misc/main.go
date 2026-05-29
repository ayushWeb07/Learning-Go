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

type ComplexNumber struct {
	a int
	b int
}

func (c *ComplexNumber) display() {
	fmt.Printf("%d + i%d", c.a, c.b)
}

func addComplexNums(c1 *ComplexNumber, c2 *ComplexNumber) *ComplexNumber {
	res := &ComplexNumber{
		a: c1.a + c2.a,
		b: c1.b + c2.b,
	}

	return res
}

func multiplyComplexNums(c1 *ComplexNumber, c2 *ComplexNumber) *ComplexNumber {
	res := &ComplexNumber{
		a: (c1.a * c2.a) - (c1.b * c2.b),
		b: (c1.a * c2.b) + (c1.b * c2.a),
	}

	return res
}

func main() {
	// create a slice
	nums := []int{5, 1, 3, 2, 4}
	fmt.Println("Before:", nums)

	InsertionSort(&nums)
	fmt.Println("After:", nums)

	// create 2 complex numbers
	c1 := &ComplexNumber{
		a: 2,
		b: 5,
	}

	c2 := &ComplexNumber{
		a: 3,
		b: 8,
	}

	fmt.Println()
	c1.display()

	fmt.Println()
	c2.display()

	// add complex nums
	addRes := addComplexNums(c1, c2)
	fmt.Println()
	fmt.Println()
	addRes.display()

	// multiply complex nums
	prodRes := multiplyComplexNums(c1, c2)
	fmt.Println()
	fmt.Println()
	prodRes.display()

}
