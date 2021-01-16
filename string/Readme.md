# String

A string is a slice of bytes in Go. Strings can be created by enclosing a set of characters inside double quotes " ".

since string is a slice of bytes it is possibe o access each byte of slice by looping theogh it

# Rune

A rune is a builtin type in Go and it's the alias of int32. Rune represents a Unicode code point in Go. It doesn't matter how many bytes the code point occupies, it can be represented by a rune.

## Creating a string from a slice of bytes

```Go
package main

import (
    "fmt"
)

func main() {
    byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice)
    fmt.Println(str)
}
```

byteSlice in of the program above contains the UTF-8 Encoded hex bytes of the string Café. The program prints <code> Café </code>

## String length

The <code>RuneCountInString(s string) (n int) </code> function of the utf8 package can be used to find the length of the string. This method takes a string as an argument and returns the number of runes in it.

## String comparison

The == operator is used to compare two strings for equality. If both the strings are equal, then the result is true else it's false.
**Gotacha** Tripple quality sign is not suppported in Go

## String concatenation

There are multiple ways to perform string concatenation in Go. Let's look at a couple of them.

The most simple way to perform string concatenation is using the + operator.

The second way to concatenate strings is using the <code>Sprintf</code> function of the fmt package.

The <code>Sprintf</code> function formats a string according to the input format specifier and returns the resulting string.

## Strings are immutable

Strings are immutable in Go. Once a string is created it's not possible to change it.

To workaround this string immutability, strings are converted to a slice of runes. Then that slice is mutated with whatever changes are needed and converted back to a new string.
