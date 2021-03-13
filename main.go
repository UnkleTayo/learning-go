package main

import (
	"log"

	"github.com/unkletayo/learninggo/channels"
	"github.com/unkletayo/learninggo/path_seperator"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := channels.RandomNumber(10)

	// Passing a value into the channel
	intChan <- randomNumber
}

func main() {
	path_seperator.Path()
	intChan := make(chan int)

	defer close(intChan)

	go CalculateValue(intChan)

	// Getting a value from the channel
	num := <-intChan
	log.Println(num)
}
