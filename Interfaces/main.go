package main

import (
	"fmt"
	"log"
)

// interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
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
func assert(i interface{}) {
	s, ok := i.(int)
	fmt.Println(s, ok)
}

func main() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())

	var s interface{} = 57
	assert(s)
	var i interface{} = "Steven Paul"
	assert(i)

	findType("Kakashi")
	findType(77)
	findType(89.98)

	dog := Dog{
		Name:  "Fiddo",
		Breed: "GS",
	}

	PrintInfo(dog)

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberOfTeeth: 43,
	}
	PrintInfo(gorilla)
}

// Type switch
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "GRUNT"
}

func (g Gorilla) NumberOfLegs() int {
	return 4
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs())
}
