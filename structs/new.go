package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Kakashi", lastName: "Hatake"}
	fmt.Println(alex)

	
var bobby person
fmt.Println(bobby)

// listing out the values inside the Go 
fmt.Printf("%+v", alex)

}