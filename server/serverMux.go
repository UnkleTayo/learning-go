package main

import (
	"io"
	"log"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	io.Writer(w, "World!")
}

func handlerMux() {
	mux := http.NewServerMux()
	mux.HandleFunc("/", greet)
	err := http.ListenAndServe(connHost+":"+connPort, handlers.CompressHandlers(mux))
	if err != nil {
		log.Fatal("error starting http server", err)
		return
	}
}
