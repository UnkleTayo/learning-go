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

*Assigning and slicing*

Assigning a value to any specific index within a slice is identical to how you do this
with arrays. To change the value of an individual element, use the [ ] operator.

```Go
// Create a slice of integers.
// Contains a length and capacity of 5 elements.
slice := []int{10, 20, 30, 40, 50}
// Change the value of index 1.
slice[1] = 25
```

Taking slice of a slice(Slicing a slice)

```Go
// Create a slice of integers.
// Contains a length and capacity of 5 elements.
slice := []int{10, 20, 30, 40, 50}
// Create a new slice.
// Contains a length of 2 and capacity of 4 elements.
newSlice := slice[1:3]
```

*How length and capacity are calculated*

```Go
For slice[i:j] with an underlying array of capacity k
Length: j - i
Capacity:k-i
```

Another way to look at this is that the first value represents the starting index position of the element the new slice will start with—in this case, 1. The second value represents the starting index position (1) plus the number of elements you want to include (2); 1 plus 2 is 3, so the second value is 3. Capacity will be the total number of elements associated with the slice.

## Consequences of  making changes to an array

```Go
// Create a slice of integers.
// Contains a length and capacity of 5 elements.
slice := []int{10, 20, 30, 40, 50}
// Create a new slice.
// Contains a length of 2 and capacity of 4 elements.
newSlice := slice[1:3]
// Change index 1 of newSlice.
// Change index 2 of the original slice.
newSlice[1] = 35
```

## GROWING SLICES

One of the advantages of using a slice over using an array is that you can grow the capacity of your slice as needed. Go takes care of all the operational details when you use the built-in function append.

To use append a source slice and a value to be appended is needed. The append calls provide a new slice with the changes. Append will increase the length of the new slice. The capacity may not necessarily be affected, it is dependent on the capacity of the source slice.

```Go
// Create a slice of integers.
// Contains a length and capacity of 5 elements.
slice := []int{10, 20, 30, 40, 50}
// Create a new slice.
// Contains a length of 2 and capacity of 4 elements.
newSlice := slice[1:3]

Allocate a new element from capacity
Assign a value of 60 to the new element
newSlice = append(newSlice, 60)
```

Because there was available capacity in the underlying array for newSlice, the
append operation incorporated the available element into the slice’s length and assigned the value. Since the original slice is sharing the underlying array, slice also sees the changes in index 3.

When there’s no available capacity in the underlying array for a slice, the append function will create a new underlying array, copy the existing values that are being referenced, and assign the new value.

## Increasing length and capacity of a slice with append

```Go
create a new slice of length and capacity 4
slice := []int{10,2,4,50,69}

Append a new slice to the slice
newSlice := append(slice, 50)

```

## Three Index Slices

This is used to control the capacity of a newly crated slice. Its purpose is not to increase the capacity but to restrict the capacity.
Being able to restrict capacity of a new slice provides protection to the underlying array and also gives control over append operation

Let’s start with a slice of five strings that contain fruit you can find in your local supermarket.

```Go
// Create a slice of strings.
// Contains a length and capacity of 5 elements.
source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

// slice the third element and increase it capacity 
slice := source[2:3:4]
```

## Calculating length and capacity

```Go
For slice[i:j:k] or [2:3:4]
Length: j - i or 3-2=1
Capacity: k - i or 4 - 2 = 2

```

Again, the first value represents the starting index position of the element the new slice will start with—in this case, 2. The second value represents the starting index position (2) plus the number of elements you want to include (1); 2 plus 1 is 3, so the second value is 3. For setting capacity, you take the starting index position of 2, plus the number of elements you want to include in the capacity (2), and you get the value of 4.


 By having the option to set the capacity of a new slice to be the same as the length, you can force the first append operation to detach the new slice from the underlying array. Detaching the new slice from its original source array makes it safe to change. 


The built-in function append is also a variadic function. This means you can pass multiple values to be appended in a single slice call. If you use the ... operator, you can append all the elements of one slice into another. 

## Appending slice from another slice

```Go
// Create two slices each initialized with two integers.
s1 := []int{1, 2}
s2 := []int{3, 4}
// Append the two slices together and display the results.
fmt.Printf("%v\n", append(s1, s2...))
Output:
[1 2 3 4]
```

## Iterating over slices
```Go
// create a slice of integers

slice := []int{10,20,40,56,67}

// iterate over

for index, value := range slice{
    fmt.Printf("Index: %d value :%d\n", index value)
}
```

The keyword <code>range</code>, when iterating over a slice, will return two values. The first value is the index position and the second value is a copy of the value in that index position.

It’s important to know that range is making a copy of the value, not returning a reference. If you use the address of the value variable as a pointer to each element, you’ll be making a mistake. 

