package main

import(
	"fmt"
	"runtime"
)

func main(){
	go foo()


	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Routine\t", runtime.NumGoroutine())

}


func foo(){
	fmt.Println("Hello")
}