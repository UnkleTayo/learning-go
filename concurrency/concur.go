package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Concur() {
	runtime.GOMAXPROCS(1)

	// wg is used to wait for the program to finish
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Routines")

	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}

	}()

	wg.Wait()

	fmt.Println("\nTerminating Program")
}
