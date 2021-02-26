package main

import (
	"fmt"
	"os"
)

func main(){
	// fmt.Println("path \n", os.Args[0])
	// fmt.Println("1st Arg \n", os.Args[1])
	// fmt.Println("2nd Arg \n", os.Args[2])
	// fmt.Println("3rd Arg", os.Args[3])

	name := os.Args[1]
	name2 := os.Args[2]
	name3 := os.Args[3]

	fmt.Println("Hello", name, name2,"and	", name3, "!")
	name, age := "Kakashi", 2018

	fmt.Println("My name is ", name)
	fmt.Println("My age is ", age)
	fmt.Println("I am the 6th Hokage")
}