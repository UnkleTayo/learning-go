package main

import (
	"fmt"
	"unsafe"
)
// The following are the basic types available in go

// bool
// Numeric Types
// int8, int16, int32, int64, int
// uint8, uint16, uint32, uint64, uint
// float32, float64
// complex64, complex128
// byte
// rune
// string

func main(){
	a:= true
	b:= false
	fmt.Println("a:", a, "b:", b)
	c := a && b
  fmt.Println("c:", c)
  d := a || b
	fmt.Println("d:", d)

	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) //type and size of a
    fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b

// 	Signed integers
// int8: represents 8 bit signed integers
// size: 8 bits
// range: -128 to 127

// int16: represents 16 bit signed integers
// size: 16 bits
// range: -32768 to 32767

// int32: represents 32 bit signed integers
// size: 32 bits
// range: -2147483648 to 2147483647

// int64: represents 64 bit signed integers
// size: 64 bits
// range: -9223372036854775808 to 9223372036854775807

// Unsigned integers
// uint8: represents 8 bit unsigned integers
// size: 8 bits
// range: 0 to 255

// uint16: represents 16 bit unsigned integers
// size: 16 bits
// range: 0 to 65535

// uint32: represents 32 bit unsigned integers
// size: 32 bits
// range: 0 to 4294967295

// uint64: represents 64 bit unsigned integers
// size: 64 bits
// range: 0 to 18446744073709551615

// uint : represents 32 or 64 bit unsigned integers depending on the underlying platform.
// size : 32 bits in 32 bit systems and 64 bits in 64 bit systems.
// range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems

// Floating point types
// float32: 32 bit floating point numbers
// float64: 64 bit floating point numbers

// The following is a simple program to illustrate integer and floating point types

// Complex types
// complex64: complex numbers which have float32 real and imaginary parts
// complex128: complex numbers with float64 real and imaginary parts

// The builtin function complex is used to construct a complex number with real and imaginary parts. The complex function has the following definition

// func complex(r, i FloatType) ComplexType  
// It takes a real and imaginary part as a parameter and returns a complex type. Both the real and imaginary parts must be of the same type. ie either float32 or float64. If both the real and imaginary parts are float32, this function returns a complex value of type complex64. If both the real and imaginary parts are of type float64, this function returns a complex value of type complex128

// Complex numbers can also be created using the shorthand syntax

// c := 6 + 7i  
// Let's write a small program to understand complex numbers.

// c1 := complex(5, 7)
// c2 := 8 + 27i
// cadd := c1 + c2
// fmt.Println("sum:", cadd)
// cmul := c1 * c2
// fmt.Println("product:", cmul)

	var f int = 89
	h := 95
	fmt.Println("value of a is", f, "and b is", h)
}