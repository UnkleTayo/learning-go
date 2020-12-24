package main

import "fmt"


type contactInfo struct {
	email string
	zip string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}


func main() {
	// alex := person{firstName: "Kakashi", lastName: "Hatake"}
	// fmt.Println(alex)

	
// var bobby person
// fmt.Println(bobby)

// listing out the values inside the Go 
// fmt.Printf("%+v", alex)

// Embedding structs 
	jim := person{
		firstName: "Jim",
		lastName: "Longstron",
		contactInfo: contactInfo {
			email: "jim@mail.com",
			zip: "30303",
		},
	}

	jim.updateName("jimmy")
	jim.print()
}
func(p *person) updateName(newFirstName string){
	p.firstName = newFirstName
}

// receiver function 
func (p person) print(){
	fmt.Printf("%+v", p)
}