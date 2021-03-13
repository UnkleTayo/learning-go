package main

import (
	"fmt"
	"time"
)

func numbers() {  
	for i := 1; i <= 5; i++ {
			time.Sleep(250 * time.Millisecond)
			fmt.Printf("%d ", i)
	}
}
func alphabets() {  
	for i := 'a'; i <= 'e'; i++ {
			time.Sleep(400 * time.Millisecond)
			fmt.Printf("%c ", i)
	}
}

func main() {

	// routines
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")

	var a chan int
	if a == nil{
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
}

// sending and receivng from channel 