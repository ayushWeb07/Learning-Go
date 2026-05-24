package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch case
	day := 5

	switch day {
	case 1:
		fmt.Println("Mon")

	case 2:
		fmt.Println("Tues")

	case 3:
		fmt.Println("Wed")

	case 4:
		fmt.Println("Thurs")

	case 5:
		fmt.Println("Fri")

	default:
		fmt.Println("Weekend")
	}

	fmt.Println()

	// using time module
	wd := time.Now().Weekday()

	switch wd {
	case time.Monday:
		fmt.Println("Mon")

	case time.Tuesday:
		fmt.Println("Tue")

	case time.Wednesday:
		fmt.Println("Wed")

	case time.Thursday:
		fmt.Println("Thu")

	case time.Friday:
		fmt.Println("Fri")

	default:
		fmt.Println("Weekend")
	}

	fmt.Println()

	// type switch
	whatsMyType := func(i interface{}) {
		switch i.(type) {
		case string:
			fmt.Println("Its a string", i)
			
		case int:

			fmt.Println("Its an int", i)

		case bool:
			fmt.Println("Its a bool", i)

		default:
			fmt.Println("Its some other type", i)
		}
	}

	whatsMyType("ayush")
	whatsMyType(19)
	whatsMyType(false)
	whatsMyType(time.Now())
}
