// Structs is adata struture that allows composition of different data types it synonymous to class in JS
package main

import (
	"fmt"
)

// Creating a value of type person
type person struct{
	first string
	last string
	age int
}

// Embeded 🇬🇸 
type secretAgent struct{
	// assigning type of person to secretAgent
	person
	ltk bool
	first string
}
// Keyword identifier type 
// func (r reciever) identifier(parmeter) (return (s)) {code}
func(s secretAgent) speak(){
	fmt.Println("I am", s.first, s.last)
}

func main(){
sa1:= secretAgent{
	person: person{
		first: "Kakashi",
		last: "Hatake",
		age: 40,
	},
	first: "Ninja",
	ltk: true,
}
	  
	p1:= person {
		first: "Bond",
		last:"Manny",
		age: 23,
	}

	p2:= person {
		first: "Kakashi",
		last:"Hatake",
		age: 45,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last) 
	fmt.Println(p2.first, p2.last)
	fmt.Println(sa1)
	sa1.speak()
	// Acessong the props inside 
	fmt.Println(sa1.first, sa1.last, sa1.person.first)

	p3:=struct{
		first string
		last string
		age int
	}{
		first: "Kakashi",
		last: "Hatake",
		age: 45,
	}

	fmt.Println(p3)
}