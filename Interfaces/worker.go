package main

import "fmt"

// Worker should have comment
type Worker interface {
	Work()
}

// SinglePerson receiver function
type SinglePerson struct {
	name string
}

//Work executed by a single Person type
func (p SinglePerson) Work() {
	fmt.Println(p.name, "is working")
}

func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func workers() {
	p := Person{
		name: "Naveen",
	}
	var w Worker = p
	describe(w)
	w.Work()
}
