package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

// Printing characters
func printChars(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func compareStrings(str1 string, str2 string) {  
	if str1 == str2 {
			fmt.Printf("%s and %s are equal\n", str1, str2)
			return
	}
	fmt.Printf("%s and %s are not equal\n", str1, str2)
}


// Mutating a string 
func mutate(s []rune) string {  
	s[0] = 'a' 
	return string(s)
}

func main() {
	name := "Señor"
	charsAndBytePosition(name)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(name))


	string1 := "Go"
	string2 := "is awesome"
	result := fmt.Sprintf("%s %s", string1, string2)
	fmt.Println(result)

	// Mutating string
	h := "hello"
    fmt.Println(mutate([]rune(h)))
}
