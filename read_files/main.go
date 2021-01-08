package main

import (
	"fmt"
	"io/ioutil"
	"os"
)


func main() {
// creating byte of slice 
mydata:= []byte("This code is sweet to write\n")


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



	// writing to an existig file 
	f, err := os.OpenFile("myfile.data", os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil{
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("This was meant to be a joke but lets see\n"); err != nil{
		panic(err)
	}

	data, err = ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

}