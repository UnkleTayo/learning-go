package main

import "fmt"

func SwitchCase() {
	finger := 5
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
case 2:
		fmt.Println("Index")
case 3:
		fmt.Println("Middle")
case 4:
		fmt.Println("Ring")
case 5:
		fmt.Println("Pinky")
default: //default case
		fmt.Println("incorrect finger number")
	}
}

func number() int {  
	num := 15 * 5 
	return num
}

func fallThrogh() {
	switch num := number(); { //num is not a constant
	case num < 50:
			fmt.Printf("%d is lesser than 50\n", num)
			fallthrough
	case num < 100:
			fmt.Printf("%d is lesser than 100\n", num)
			fallthrough
	case num < 200:
			fmt.Printf("%d is lesser than 200", num)
	}
}