package main

import (
	"encoding/json"
	"sort"
	"fmt"
)

type person struct{
	First string `json:"First"`
	Last string  `json:"Last"`
	Age int  `json:"Age"`
}

func main(){
	s := `[{"First":"Kakashi","Last":"Hatake","Age":42},{"First":"Naruto","Last":"Uzumaki","Age":24}]`

	a:= []int{1,4,5,6,9,4,19,6,}
	sort.Ints(a)
	// b:= sort.Reverse(a)

	fmt.Println(a)
	// fmt.Println(b)

	bs :=[]byte(s)

	var people []person 

	err := json.Unmarshal(bs, &people)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)
	
	for i, v :=range people{
		fmt.Println("\n Person Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}