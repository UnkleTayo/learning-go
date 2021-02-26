package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	msg := os.Args[1]
	
	l := len(msg)
	h := l

	s := strings.Repeat("!", h) + msg + strings.Repeat("!",  l);
	s = strings.ToUpper(s)

	fmt.Println(s)
}