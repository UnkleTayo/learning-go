package main

import "fmt"

// interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

// Type assertion 
func assert( i interface{}){
	s, ok := i.(int)
	fmt.Println(s, ok)
}


func main() {
name := MyString("Sam Anderson")
var v VowelsFinder
v = name
fmt.Printf("Vowels are %c", v.FindVowels())

var s interface{} =57
assert(s)
var i interface{} = "Steven Paul"
assert(i)

findType("Kakashi")
findType(77)
findType(89.98)
}

// Type switch 
func findType(i interface{}){
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
default:
		fmt.Printf("Unknown type\n")
}
}