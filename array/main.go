// creating an array
// var name [size] type
package main

import (
	"fmt"
)

func main(){
// declaring an array with elements
days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

fmt.Println(days[0]) // prints 'monday'
fmt.Println(days[5]) // prints 'saturday'

}