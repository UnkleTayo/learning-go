package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

// handling routes
func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting http server", err)
		return
	}
}

// func main() {
// http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
// 	fmt.Println(writer, "Hello World")
// })

// server := http.Server{
// 	Addr: "9001",
// }
// fmt.Println("starting server")

// if err := server.ListenAndServe(); err != nil {
// 	panic(err)
// }

// Handler()
// Handling()
// }
