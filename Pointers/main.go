package main

import "fmt"


func hello() *int {  
	i := 5
	return &i
}

func main() {
	b := 255
    a := &b
    fmt.Println("address of b is", a)
    fmt.Println("value of b is", *a)
    *a++
    fmt.Println("new value of b is", b)

	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)

	d := hello()
	fmt.Println("Value of d", *d)

}