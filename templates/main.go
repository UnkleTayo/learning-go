package main

import (
	"os"
	"text/template"
)


func main() {

	// parse the file  using template.parsefiles
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data:= struct{
		Name string
	}{"John Snow"}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}