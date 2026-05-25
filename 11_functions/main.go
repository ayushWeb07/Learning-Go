package main

import "fmt"

func helloWorld() {
	fmt.Println("Hello World")
}

func add(a int, b int) int {
	return a + b
}

func prod(x, y, z int) int {
	return x * y * z
}

func getNameAndWins() (string, int) {
	return "Khabib", 29
}

func printSquareAndCube(sq func(x int) int, cube func(x int) int, x int) {
	sqRes := sq(x)
	cubeRes := cube(x)

	fmt.Println("\n", x, "*", x, "=", sqRes)
	fmt.Println(x, "*", x, "*", x, "=", cubeRes)
}

func positivityDetector() func(x int) string {
	return func(x int) string {
		if x > 0 {
			return "+ve"
		} else if x < 0 {
			return "-ve"
		} else {
			return "zero"
		}
	}
}

func main() {
	// simple function without any params
	helloWorld()
	helloWorld()

	// function with params
	sumRes := add(6, 3)
	fmt.Println("\n6 + 3 ->", sumRes)

	prodRes := prod(2, 5, 3)
	fmt.Println("\n2 * 5 * 3 ->", prodRes)

	// function with multiple params
	name, wins := getNameAndWins()
	fmt.Println("\n", name, wins)

	// pass func as args
	sq := func(x int) int {
		return x * x
	}

	cube := func(x int) int {
		return x * x * x
	}

	printSquareAndCube(sq, cube, 7)

	// get func as output
	f := positivityDetector()
	r1 := f(2)
	r2 := f(-1)
	r3 := f(0)

	fmt.Println("\n", r1)
	fmt.Println(r2)
	fmt.Println(r3)

}
