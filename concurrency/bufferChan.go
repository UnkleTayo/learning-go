package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {  
	for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}


func main() {  

		ch := make(chan int, 2)
    go write(ch)
    time.Sleep(2 * time.Second)
    for v := range ch {
        fmt.Println("read value", v,"from ch")
        time.Sleep(2 * time.Second)

		}
		
		cdo := make(chan string, 3)
    cdo <- "naveen"
    cdo <- "paul"
    fmt.Println("capacity is", cap(cdo))
    fmt.Println("length is", len(cdo))
    fmt.Println("read value", <-cdo)
    fmt.Println("new length is", len(cdo))
}