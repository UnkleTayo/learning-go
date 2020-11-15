package main

import "fmt"

// Mutiple constants can be wrapped in paranthesis, likewise packages being imported 
const (
	pi       = 3.14
	language = "Go"
)

func main() {
	fmt.Println(pi)
	fmt.Println(language)
}
