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

fmt.Println("Array Pointer ", *arrPointer[0])

// Declare a string pointer array od three elements
var arr1 [3]*string


// Declare a second string pointer array of three elements.
// Initialize the array with string pointers.
arr2 := [3]*string{new(string), new(string), new(string)}

*arr2[0]="Red"
*arr2[1] = "Blue"
*arr2[2] = "Green"
// Copy the values from array2 into array1.
arr1 = arr2

fmt.Println("Array 1", arr1)
fmt.Println("Array 2", arr2)

// Mutidimensional Array 
 var arr3 [4][2]int

 arr4 := [4][2]int {{10, 13}, {14, 15}, {14, 98}, {30, 24}}

fmt.Println("Array 3", arr3)
fmt.Println("Array 4", arr4)
}