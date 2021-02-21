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
