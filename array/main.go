// creating an array
// var name [size] type
package main

import (
	"fmt"
)

func main(){
	// Declare a string array.
	// Initialize each element with a specific value.
	// Capacity is determined based on the number of values initialized.

days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

fmt.Println(days[0]) // prints 'monday'
fmt.Println(days[5]) // prints 'saturday'

// Declare an integer array of five elements.
// Initialize index 1 and 2 with specific values.
// The rest of the elements contain their zero value.
array := [5]int{1: 10, 2: 20}
fmt.Println(array)

// Declare an integer pointer array of five elements.
// Initialize index 0 and 1 of the array with integer pointers.
arrPointer := [5]*int{0: new(int), 1: new(int)}

// Assign values to index 0 and 1.
*arrPointer[0] = 10
*arrPointer[1] = 20

}