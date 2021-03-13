package main

import (
	"log"
)

func main() {
	nested()
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			log.Printf("%d", i)
		}
	}
}

// break
func nested() {
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			log.Println("*")
		}
	}
}
