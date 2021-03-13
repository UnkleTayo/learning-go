package main

import "fmt"

// This comment should make it stop screeming
type Employee struct {
	name string
	age  int
}

// creating type alias
type myInt int

// Method with value receiver
func (e Employee) changeName(newName string) {
	e.name = newName
}

// Method with pointer receiver
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}

	p := &e
	fmt.Printf("Type of %x\n", p)

	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)

	letsSee()
}

// DEFINING METHID ON STRUCT TYPE
// In the program above, in line no. 3 we are trying to add a method named add on the built-in type int. This is not allowed since the definition of the method add and the definition of type int is not in the same package.
func (a myInt) add(b myInt) myInt {
	return a + b
}
