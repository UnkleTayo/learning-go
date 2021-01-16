// Everything in Go is always about TYPE 
// You declare a type of everything(int, string struct)
// Just like Javascript you can print to your termainal using the fmt package 

// There are different (methods) attached to fmt, some I am famoliar with at this stage are Print, Println, Sprint, Sprintln

// Now lets get to the task of this module
// I will be printing a number of type int in differnt bases 
// In Go  base 10, 2, and 16 are denoted by d, b and x respectively

package main

import "fmt"

func main(){

	// I can declare varaibles using var or through short declarations := 
	x := 24
	fmt.Printf("%d - %b - %#x", x, x, x,)
}