# Arrays and Slices

## Array

An array is a collection of elements that belong to the same type.
Go does not allow mixing of array that belongs to different types eg string and inegars

Slices in Go are just like arrays but with super powers this is because they allow you to access a subset of array's element

_Declaration_
An array belongs to type [n]T. n denotes the number of elements in an array and T represents the type of each element.
The number of elements n is also a part of the type

Slices are made up of three things, a pointer, a length, and a capacity.

```Go
package main

import (
    "fmt"
)


func main() {
    var a [3]int //int array with length 3
    fmt.Println(a) // [0,0,0]
}
```

The above declared array declares an integar array of length 3.
The elements in the array are assigned the zero value of the array type.
in this case a is assigned to 0 whicj is the zero value of int.

Index of an array starts at 0 and ends at <i>length -1<i/>

## Creating array using shorthand method

```Go
...

func main() {
    b := [3]int{12, 78, 50} // short hand declaration to create array
    fmt.Println(b)
}
```

it is not necessary that all values have to be assigned during a short declaration

```Go
...

func main() {
    b := [3]int{90} // short hand declaration to create array
    fmt.Println(b) // [90 0 0]
}
```

The length of the array can be ignored during declaration and replaced with ... and let the compiler find the length and replace it.

```Go
...

func main() {
    a := [...]int{12, 78, 50} // ... makes the compiler determine the length
    fmt.Println(a)
}
```

<strong>The size of the array is a part of the type</strong>. Hence [5]int and [10]int are distinct types.

## Arrays are value types

Arrays in Go are value types this means when they are assigned to a new variable, a copy of the original array is assigned to the new variable. If changes are made to the new variable, it will not be reflected in the original array.

Also, when arrays are passed as function parameters, they are passed by value and the original array is unchnaged

## Length of an array

The length of the array is found by passing the array as parameter to the len function.

```Go
...

func main() {
    a := [...]float64{67.7, 89.8, 21, 78}
    fmt.Println("length of a is",len(a)) // length of a is 4

}
```

## Iterating arrays using range

The for loop can be used to iterate over elements of an array.
Range can also be used

# Slices

A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are the just references to existing arrays.

## Creating a slice

A slice with elements of type T is represented by []T

The syntax a[start:end] creates a slice from array a starting from index start to index end - 1. The code of the above program a[1:4] creates a slice representation of the array a starting from indexes 1 through 3.

## Modifying a slice

A slice does not own any data of its own. It is just a representation of the underlying array. Any modifications done to the slice will be reflected in the underlying array.
