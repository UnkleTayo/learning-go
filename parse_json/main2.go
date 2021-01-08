package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


func main2() {
	// Open our jsonFile
	jsonFile, err :=  os.Open("users.json")

	// if we open Os.open it returns an err so we have to handle it 
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	bytevalue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(bytevalue), &result)


	fmt.Println(result["users"])
}