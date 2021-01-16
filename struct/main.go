package main

import "fmt"


type contactInfo struct {
	email string
	zipCode int
}

// step 1; define struct fields
type person struct {
	firstName string
	lastName string
	contactInfo
}


// create a new value of that struct type 

func main(){
// alex:= person {firstName: "Alex", lastName: "Abey"}

// var alexis person
// alexis.firstName = "hey"
// alexis.lastName="KOOK"

// Embedding structs 

// fmt.Printf("%+v", alexis)

		jim:= person {
			firstName: 	"Jim",
			lastName: "Jake",
			contactInfo: contactInfo{
				email:	 "jim@mail.com",
				zipCode:  98787,
			},
		}

// This wont work because go makes a copy of jim and didnt update it 
	// &  returns value of the memory address
	// address to value = *address
	// value to address = &value

		// jimPointer := &jim
		jim.updateName("Auwal")
		jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string){
	// * returns value of the memory address
	(*pointerToPerson).firstName = newFirstName
}

// Go shortcut 

func (p person) print(){
	fmt.Printf("%+v",p)
}