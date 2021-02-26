package main

import (
	"fmt"
)

func main(){
// Just like the previous section where I mentioned that the values of number in different bases can be gotten. Go also makes it easy to get the UTF values of numbers with the use of  %q 

// In the code snippet below I looped through the numbers between 60 and 122 to get the corresponding UTF 8 values

for i := 60; i < 122; i++ {
	fmt.Println("number \t base2 \t hexadecimal\t UTF character")
	fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
}

}