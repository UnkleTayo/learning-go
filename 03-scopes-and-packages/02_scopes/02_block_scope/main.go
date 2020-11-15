package main

// This is a file scope
import "fmt"


func info() { // block scope starts

	var name string = "Kakashi Hatake"
	var ok  bool= true

	// hello and ok are visible here but name cant be acessed from outside the main function
	fmt.Println(name, ok)

} // block scope ends


func main() { // block scope starts



	// hello and name are not visible here 

	// This will return an error has they cant be acceed from within this function
	// fmt.Println(name, ok)

} // block scope ends
