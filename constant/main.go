// Declaring a constant
// The keyword const is used to declare a constant. Let's see how to declare a constant using an example.

package main

import (  
    "fmt"
)

func main() {  
		const (
			name = "John"
			age = 50
			a = 50
			country = "Canada"
		)
		fmt.Println(a)
		
		// Declaring a group of constants
// There is also another syntax to define a group of constants using a single statement. An example to define a group of constants using this syntax is provided below.


fmt.Println(name)
fmt.Println(age)
fmt.Println(country)
}

// The value of a constant should be known at compile time. Hence it cannot be assigned to a value returned by a function call since the function call takes place at run time.