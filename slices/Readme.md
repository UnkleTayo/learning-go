# Slices

A slice is a data structure that provides a way for you to work with and manage collections of data.

A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are the just references to existing arrays.

Slices are built around the concept of dynamic arrays that can grow and
shrink as you see fit. They’re flexible in terms of growth because they have their own built-in function called <i>append</i>, which can grow a slice quickly with efficiency.

Slices in Go are just like arrays but with super powers this is because they allow you to access a subset of array's element

## Internals

Slices are tiny objects that abstract and manipulate an underlying array.
Slices are three field data structures containing metadata needed by Go

The three fields are a pointer to the underlying array, are

-A pointer to the underlying array,
-The length or the number of elements the slice has access to
-The capacity or the number of elements the slice has available for growth.

## Creating and initializing

There are several ways to create and initialize slices in Go. Knowing the capacity you need ahead of time will usually determine how you go about creating your slice.

*Make and Slice Literals*

## Creating a slice

One way to create a slice is to use the built-in function make. When you use make, one option you have is to specify the length of the slice.

```Go
... 

func main() {
   // Create a slice of strings.
    // Contains a length and capacity of 5 elements.
    slice := make([]string, 5)
}

```

When you just specify the length, the capacity of the slice is the same. You can also specify the length and capacity separately

```Go
... 

func main() {
   // Create a slice of strings.
    // Contains a length and capacity of 5 elements.
    slice := make([]int 3, 5)
}

```

 Trying to create a slice with a capacity that’s smaller than the length is not allowed

## Declaring a slice with slice literal

```Go
... 

func main() {
// Create a slice of strings.
// Contains a length and capacity of 5 elements.
slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}

// Create a slice of integers.
// Contains a length and capacity of 3 elements.
slice2 := []int{10, 20, 30}
}

```

## Declaring a slice with index position

When using a slice literal, you can set the initial length and capacity. All you need to do is initialize the index that represents the length and capacity you need.

```Go
... 

func main() {
// Create a slice of strings.
// Initialize the 100th element with an empty string.
slice := []string{99: ""}
}

```

## Difference between arrays and slice

If you specify a value inside the [ ] operator, you’re creating an array. If
you don’t specify a value, you’re creating a slice.

```Go
// Create an array of three integers.
array := [3]int{10, 20, 30}

// Create a slice of integers with a length and capacity of three.
slice := []int{10, 20, 30}
```

## NIL AND EMPTY SLICES

There are times when we need to declare a nil slice. A nil slice is created
by declaring a slice without any initialization.

```Go
// Create a nil slice of integers.
var slice []int // nil slice, length = 0, capacity = 0
```

A nil slice is the most common way you create slices in Go. They can be used with many of the standard library and built-in functions that work with slices. 
They’re useful when you want to represent a slice that doesn’t exist, such as when an exception occurs in a function that returns a slice

## Declaring an empty slice
```Go
// Use make to create an empty slice of integers.
slice := make([]int, 0)
// Use a slice literal to create an empty slice of integers.
slice := []int{}
```

## Working with slices

