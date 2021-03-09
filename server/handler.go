package main

import (
	"fmt"
	"net/http"
)

type customHandler struct {
	response string
}

func (c *customHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, c.response)
}

// Handler Function
func Handler() {
	server := http.Server{
		Addr:    ":9091",
		Handler: &customHandler{"Hello world from the custom handler"},
	}
	fmt.Println("Starting server")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
