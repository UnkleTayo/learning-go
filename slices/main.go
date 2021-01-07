package main

import (
	"fmt"
)

func main() {
	x:= make([]int, 10,100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x= append(x, 4555)
	fmt.Println(x)


	youtubeSubscribers := map[string]int{
		"TutorialEdge":     2240,
		"MKBHD":            6580350,
		"Fun Fun Function": 171220,
	}
	
	fmt.Println(youtubeSubscribers["MKBHD"]) 
}