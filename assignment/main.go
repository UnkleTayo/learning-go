package main

import (
	"fmt"
	"strconv"
	"time"
)

func main(){
	var(
		speed int
		now time.Time
	)

	eq:= "1 + 3 = "
	sum := 1 + 2
	rata, _ := strconv.Atoi("5")
	fmt.Println(rata)

	fmt.Println(eq + strconv.Itoa(sum))
speed, now = 100, time.Now()
	fmt.Println(speed, now)
}