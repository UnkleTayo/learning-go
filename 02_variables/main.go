package main

import "fmt"

// or making a function call.

func main() {
	// single variable declaration with initial value
	var name string = "Kakashi Hatake"
	// Multiple variable declaration  
	var width, height int = 100, 50

// The short declaration := can only be used withing a function and not outside 
	a := 10
	b := "golang"
	c := 4.17
	


	// This will print the type and value of the variables declared above 
	fmt.Println("width is", width, "height is", height)
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", name)


	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", name)
}
