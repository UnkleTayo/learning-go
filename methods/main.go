package main

import "fmt"

// This comment should make it stop screeming
type Employee struct {
	Name string
}

// create a reference function
func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

// create a reference function
func (e *Employee) PrintName() {
	fmt.Println(e.Name)
}

func main() {
	var employee Employee
	employee.Name = "Elliot"
	employee.UpdateName("Forbsey")
	employee.PrintName()
}
