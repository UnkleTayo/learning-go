package main

import "fmt"

// A CONSTANT is a simple unchanging value

const bankers = "death & taxes"
// Can also be Mutiple
const (
	pi       = 3.14
	language = "Go"
)

func main() {

	const q = 42

	fmt.Println("bankers - ", bankers)
	fmt.Println("q - ", q)
	fmt.Println(pi)
	fmt.Println(language)
}


