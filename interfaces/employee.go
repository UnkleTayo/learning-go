package main

import "fmt"

type Employee interface {
	Language() string
	Age() int
	Random() (string, error)
}

type Engineer struct {
	Name      string
	age       int
	randomVal string
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
	return e.age
}

func (e Engineer) Random() {
return	res, err := fmt.Println(res)

	if err != nil {
		panic("Error Occurred")
	}
}

func employee() {
	// This will throw an error
	var programmers []Employee
	elliot := Engineer{
		Name: "Elliot",
		age:  34,
		randomVal: "Hi"
	}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, elliot)
}