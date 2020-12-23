package main

import (
	"fmt"
	"learnpackage/simpleinterest"
	"log"
)

var p,r,t = 5000.0, 10.0, 1.0

func init(){
	println("Main Package Initialized")

	if p< 0{
		log.Fatal("Pricipal is less than zero")
	}
	if r< 0{
		log.Fatal("Rate is less than zero")
	}
	if t< 0{
		log.Fatal("Duration is less than zero")
	}
}

func main() {  
		fmt.Println("Simple interest calculation in Main")
    si := simpleinterest.Calculate(p, r, t)
    fmt.Println("Simple interest is", si)
}