package main

import "fmt"


type contactInfo struct {
	email string
	zip string
}

type person struct {
	firstName string
	lastName  string
	contactInfo contactInfo
}


func main() {
	alex := person{firstName: "Kakashi", lastName: "Hatake"}
	fmt.Println(alex)

	
var bobby person
fmt.Println(bobby)

// listing out the values inside the Go 
fmt.Printf("%+v", alex)

// Embedding structs 
	jim := person{
		firstName: "Jim",
		lastName: "Longstron",
		contactInfo: contactInfo {
			email: "jim@mail.com",
			zip: "30303",
		},
	}

	fmt.Printf("%+v", jim)
}