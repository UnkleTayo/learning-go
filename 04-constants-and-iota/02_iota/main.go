package main

import "fmt"


// Iota kinda like set the value of the first decalred variable to zero
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

// It can also be declared by making only the fist and Iota only
const (
	d = iota // 0
	e        // 1
	f        // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
