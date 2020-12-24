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

// Embedding structs 
	jim := person{
		firstName: "Jim",
		lastName: "Longstron",
		contactInfo: contactInfo {
			email: "jim@mail.com",
			zip: "30303",
		},
	}

	// getting addrss in memory 
	// jimPointer := &jim
	// & is an operator that returns the memory adress of the value that the jim variable is pointing at 
	// fmt.Println(jimPointer)
	jim.updateName("jimmy")
	jim.print()
}

// * returns the value that the memory adderess is pointing at 
// * in front of type means the updatename can only be called with reference to a pointer of that type 
func(pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

// receiver function  

func (p person) print(){
	fmt.Printf("%+v", p)
}