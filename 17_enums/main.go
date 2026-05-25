package main

import "fmt"

// integer enums
type OrderStatus int

const (
	Pending OrderStatus = iota
	Confirmed
	Cancelled
	Delivered
)

func checkOrderStatus(o OrderStatus) {
	fmt.Println("Status:", o)
}

// string enums
type UserRole string

const (
	Admin  UserRole = "admin"
	Viewer UserRole = "viewer"
	Editor UserRole = "editor"
)

func checkUserRole(u UserRole) {
	fmt.Println("Status:", u)
}

func main() {
	checkOrderStatus(Pending)
	checkOrderStatus(Confirmed)
	checkOrderStatus(Cancelled)
	checkOrderStatus(Delivered)

	fmt.Println()
	checkUserRole(Admin)
	checkUserRole(Viewer)
	checkUserRole(Editor)
}
