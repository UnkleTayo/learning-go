package main

import "fmt"

// Implementing interfaces using pointer receivers vs value receivers

// Describer should also have comment
type Describer interface {
	Describe()
}

// Person should have comment
type Person struct {
	name string
	age  int
}

// Describe Implemented by Person
func (p Person) Describe() {
	//implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

// Address struct
type Address struct {
	state   string
	country string
}

// Describe also implments address
func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func pReceiver() {
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}

	d2 = &a
	d2.Describe()
}
