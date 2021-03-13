package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/unkletayo/learninggo/channels"
)

const numPool = 10
const portNumber = ":8080"

func CalculateValue(intChan chan int) {
	randomNumber := channels.RandomNumber(10)

	// Passing a value into the channel
	intChan <- randomNumber
}

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(3, 5)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 3 + 5 is %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32
	x = 100.0
	y = 0
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

func addValues(x, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	result := x / y

	return result, nil
}

func main() {
	// path_seperator.Path()
	// intChan := make(chan int)

	// defer close(intChan)

	// go CalculateValue(intChan)

	// // Getting a value from the channel
	// num := <-intChan
	// log.Println(num)

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
