package main

// This is a file scope
import "fmt"

// package scope i.e it can be acessed anywhere from withing this file
var ok  bool= true

// package scope
func main() { // block scope starts

	var name string = "Kakashi Hatake"

	// hello and name are visible here but name cant be acessed from outside the main function
	fmt.Println(name, ok)

} // block scope ends
