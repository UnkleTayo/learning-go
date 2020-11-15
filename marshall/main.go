package main

import(
	"encoding/json"
	"fmt"
)

type person struct{
	First string
	Last string
	Age int
}

func main(){
	p1:=  person{
		First: "Kakashi",
		Last: "Hatake",
		Age: 42,
	}

	p2:=  person{
		First: "Naruto",
		Last: "Uzumaki",
		Age: 24,
	}

	people:= []person{
		p1,
		p2,
	}


	fmt.Println(people)
	bs, err := json.Marshal(people)
if err !=nil {
	fmt.Println(err)
}
fmt.Println(string(bs))
}

