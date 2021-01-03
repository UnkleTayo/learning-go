package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter,  request *http.Request){
		fmt.Println(writer, "Hello World")
	})

	server := http.Server{
		Addr : "9001",
	}
	fmt.Println("starting server")

	if err := server.ListenAndServe(); err != nil{
		panic(err)
	}
}