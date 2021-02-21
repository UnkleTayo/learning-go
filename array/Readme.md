# Arrays

An array in Go is a fixed-length data type that contains a contiguous block of elements of the <b>same</b> type.

This could be a built-in type such as integers and strings, or it can be a struct type.

 They are valuable data structures because the memory is allocated sequentially.
 Having memory in a contiguous form can help to keep the memory you use stay loaded within CPU caches longer.

## Declaring and initializing

An array is declared by specifying the type of data to be stored and the total number of elements required, also known as the array’s length.

```go
// declaring an empty array of strings

var days []string

// declaring an array with elements

days := [...]string{"monday", "tuesday", "wednesday","thursday", "friday","saturday", "sunday"}
```

Once an array is declared, neither the type of data being stored nor its length can be changed. If you need more elements, you need to create a new array with the length needed and then copy the values from one array to the other.

When an array is initialized, each individual element that belongs to the array is initialized to its zero value.

An easy way t0 create and initialize and array is through the use of literals. Array literals allow you to declare the number of elements you need and specify value for the elements

## Accessing array pointer elements

An array is a value in Go. This means you can use it in an assignment operation. The variable name denotes the entire array and, therefore, an array can be assigned to other arrays of the same type.

## Multidimensional arrays

Arrays are always one-dimensional, but they can be composed to create multidimensional arrays. Multidimensional arrays come in handy when you need to manage data that may have parent/child relationships or is associated with a coordinate system.

```go
// Declare a two dimensional integer array of three elements
// by two elements.
var array [3][2]int

// Use an array literal to declare and initialize a two
// dimensional integer array.
array := [3][2]int{{10, 11}, {20, 21},  {40, 41}}

// Declare and initialize index 1 and 2 of the outer array.
array := [3][2]int{1: {20, 21}, 2: {40, 41}}
// Declare and initialize individual elements of the outer
// and inner array.
array := [3][2]int{1: {0: 20}, 3: {1: 41}}
```
## Passing array between functions
Passing an array between functions can be an expensive operation in terms of memory and performance. When you pass variables between functions they’re always passed by value. When your variable is an array, this means the entire array, regardless of its size, is copied and passed to the function. 