package main

import (
	"fmt"
)

func myfunction(firstName string, lastName string) (string, error) {
  return firstName + " " + lastName, nil
}

func main() {

  welcome := []string{"hello", "world"}
  change(welcome...)
  fmt.Println(welcome)
}


// Variadic function
// creating a variadic function 
func find(num int, nums ...int){
  fmt.Printf("type of nums is %T\n", nums)
  found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}

func change(s ...string) {  
  s[0] = "Go"
  s = append(s, "playground")
  fmt.Println(s)
}
