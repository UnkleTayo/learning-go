package main

import (
	"fmt"
	"io/ioutil"
)


func main() {
// creating byte of slice 
mydata:= []byte("This code is sweet to write")


// takes three permission
// file name with location, data to be written and the file mode and permission rule 
err := ioutil.WriteFile("myfile.data", mydata, 0777)

if err != nil{
	fmt.Println(err)
}

	// read the content of a local file
	data, err := ioutil.ReadFile("localfile.data")

	if err != nil {
		fmt.Println(err)
	}

// convert it to string 
	fmt.Print(string(data))
}