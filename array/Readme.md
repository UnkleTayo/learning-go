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

An easy way to create and initialize and array is through the use of literals. Array literals allow you to declare the number of elements you need and specify value for the elements

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

The length of the array can be ignored during declaration and replaced with  <code>...</code> and let the compiler find the length and replace it.

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

## Passing an array between function 

Passing an array between functions can be an expensive operation in terms of memory and performance. When you pass variables between functions, they’re always passed by value. When your variable is an array, this means the entire array, regardless of its size, is copied and passed to the function. 

 To see this in action, let’s create an array of one million elements of type int. On a 64-bit architecture, this would require eight million bytes, or eight megabytes, of memory. What happens when you declare an array of that size and pass it to a function?

 ```Go
// Declare an array of 8 megabytes.
var array [1e6]int
// Pass the array to the function foo.
foo(array)
// Function foo accepts an array of one million integers.
func foo(array [1e6]int) {
return array
}
 ```

 }
Every time the function foo is called, eight megabytes of memory has to be allocated on the stack. Then the value of the array, all eight megabytes of memory, has to be copied into that allocation. Go can handle this copy operation, but there’s a better and more efficient way of doing this. You can pass a pointer to the array and only copy eight bytes, instead of eight megabytes of memory on the stack.
