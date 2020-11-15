package main

import "fmt"

func main(){
	fmt.Println("Hello Gopher")
	// Uncomment the function below to see the error
	// panda()

	// Notice that I was able to access the panda and rancoon functions from inside the main.go file

	// This is possible because I can acess function from other files which are in the same package. This means that panda.go, racoon.go and main.go are in the same main package
}