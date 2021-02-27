package main

import "fmt"

type user struct {
	name  string
	email string
}

// Notify implements a method with a value receiver
func (u user) notify() {
	fmt.Printf("Sending user Email to %s<%s>\n", u.name, u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

// main is the entry point for the application.
func letsSee() {
	// Values of type user can be used to call methods
	//  declared with a value receiver
	bill := user{"Bill", "bill@mail.com"}
	bill.notify()

	// Pointers if type user can be used to call methods
	// declared with a value recover
	lisa := &user{"Lisa", "Lisa@example.com"}
	lisa.notify()

	// Pointers if type user can be used to call methods
	// declared with a value recover
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// Pointers if type user can be used to call methods
	// declared with a value recover
	lisa.changeEmail("Lisa@newdomain.com")
	lisa.notify()

}
