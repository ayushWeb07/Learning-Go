package main

import "fmt"

type City struct {
	cityName string
}

type State struct {
	stateName string
}

type Address struct {
	pincode string
	City
	State
}

func main() {
	adr := Address{
		pincode: "700083",
		City: City{
			cityName: "Kolkata",
		},
		State: State{
			stateName: "WB",
		},
	}

	fmt.Println(adr.pincode)
	fmt.Println(adr.cityName)
	fmt.Println(adr.stateName)
}
