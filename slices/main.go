package main

import(
	"fmt"
)

func main() {
	// x:= type{values} //compoaite literal
	// x := []int{4,5,6,6}
	// fmt.Println(x)
	// // looping over values in a slice usin range 
	// for i, v := range x {
	// 	fmt.Println(i,v)
	// }

	// Slicing a 🍕 
	// fmt.Println(x[1:3])
	// // Append to a slice 
	// x= append(x,10,11,12)
	// y:= []int{19,32,23}
	// x =append(x, y...)
	// fmt.Println(x)


	// Making a slice 
	// x:=  make([]type, length, capacity)
	x:= make([]int, 10,100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x= append(x, 4555)
	fmt.Println(x)


}