package main

import (
	"Assignment1/manager"
	"fmt"
)

func main() {
	// Example
	var _manager = manager.NewManager("Middle manager",
		400000,
		"Rozybakieva 20",
	)
	fmt.Println("Before Promotion")
	fmt.Println(_manager.GetPosition(), _manager.GetSalary(), _manager.GetAddress())
	// After promotion
	_manager.SetPosition("Top manager")
	_manager.SetSalary(600000)
	fmt.Println("After Promotion")
	fmt.Println(_manager.GetPosition(), _manager.GetSalary(), _manager.GetAddress())
}
