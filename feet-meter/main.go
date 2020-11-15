package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	arg := os.Args[1]
	arg2 := os.Args[2]

	feet,_ := strconv.ParseFloat(arg, 64)
	celcius,_ := strconv.ParseFloat(arg2, 64)

	meteres := feet * 0.3048
	farenheit := (celcius * (9/5)) + 32

	fmt.Printf("%g feet is %g in meters.\n", feet, meteres)
	fmt.Printf("%g celcius is %g in farenheit.\n", celcius, farenheit)
}