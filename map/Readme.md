# Map

A map is a data structure that provides you with an unordered collection of key/value pairs.

 You store values into the map based on a key.  The strength of a map is its ability to retrieve data quickly based on the key. A key works like an index, pointing to the value you associate with that key.

Maps are just like Go's representation of hash tables data structure, it allows you to map one arbitrary type to another.

## Internals

Just like arrays and slices, maps are also collections and can be iterated over only that they are <code>unordered</code> collections and there is no way to predict which key/value pairs will e returned.

Even when key/value pairs ate stored in same order every iteration over a map could return a different other this is because a map is implemented using hash tables.

The map’s hash table contains a collection of buckets. When you’re storing, removing, or looking up a key/value pair, everything starts with selecting a bucket. This is performed by passing the key—specified in your map operation—to the map’s hash function. The purpose of the hash function is to generate an index that evenly distributes key/value pairs across all available buckets.

The better the distribution, the quicker you can find your key/value pairs as the map grows. If you store 10,000 items in your map, you don’t want to ever look at10,000 key/value pairs to find the one you want. You want to look at the least number of key/value pairs possible. Looking at only 8 key/value pairs in a map of 10,000 items is a good and balanced map. A balanced list of key/value pairs across the right number of buckets makes this possible.

## Creating and Initializing maps

There are several ways you can create and initialize maps in Go. You can use the builtin function make, or you can use a map literal.

### Declaring a map using make

```Go
// Create a map with a key of type string and a value of type int.
colors := make(map[type of key]type of value)
```

E.g

```Go
 // Create a map with a key and value of type string.
// Initialize the map with 2 key/value pairs.
colors := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
```

Using a map literal is the idiomatic way of creating a map. The initial length will be based on the number of key/value pairs you specify during initialization.

## Assigning values to a map

```Go
// Add the Red color code to the map.
colors["Red"] = "#da1337"
```

You can create a nil map by declaring a map without any initialization. A nil map can’t be used to store key/value pairs. Trying will produce a runtime error

## Deleting items from a map

delete(map, key) is the syntax to delete key from a map. The delete function does not return any value.

```Go
delete(map, key)
```

## Length of the map

Length of the map can be determined using the len function.

## Checking if a key exists

To check if a key exist in a map the syntax is used

```Go
 value, exists := map[key]
```

Testing if a map key exists is an important part of working with maps. It allows you to write logic that can determine if you’ve performed an operation or if you’ve cached some particular data in the map. It can also be used to compare two maps to identify what key/value pairs match or are missing.

When retrieving a value from a map, you have two choices. You can retrieve the
value and a flag that explicitly lets you know if the key exists.

### Retrieving a value from a map and testing existence

```Go
// Retrieve the value for the key "Blue".
value, exists := colors["Blue"]
// Did this key exist?
if exists {
fmt.Println(value)
}
```

## Maps are reference types

Similar to slices, maps are reference types. When a map is assigned to a new variable, they both point to the same internal data structure. Hence changes made in one will reflect in the other.

## Maps Equality

Maps can't be compared using the == operator. The == can be only used to check if a map is nil.

One way to check whether two maps are equal is to compare each one's individual elements one by one.

## Map Lookup

When looking up a key in a map, you have two options: you can assign a single variable or two variables for the lookup call.
 The first variable is always the value returned for the key lookup, and the second value, if specified, is a Boolean flag that reports whether the key exists or not. When a key doesn’t exist, the map will return the zero value for the type of value being stored in the map. When the key does exist, the map will return a copy of the value for that key.

## Passing maps between functions

Passing a map between two functions doesn’t make a copy of the map. In fact, you can pass a map to a function and make changes to the map, and the changes will be reflected by all references to the map.

### Summary on Array, Slices and Map

- Arrays are the building blocks for both slices and maps.
- Slices are the idiomatic way in Go you work with collections of data. Maps are the way you work with key/value pairs of data.
- The built-in function make allows you to create slices and maps with initial
length and capacity. Slice and map literals can be used as well and support setting initial values for use.
- Slices have a capacity restriction, but can be extended using the built-in function append.
- Maps don’t have a capacity or any restriction on growth.
-The built-in function len can be used to retrieve the length of a slice or map.
- The built-in function cap only works on slices.
- Through the use of composition, you can create multidimensional arrays and
slices. You can also create maps with values that are slices and other maps. A
slice can’t be used as a map key.
- Passing a slice or map to a function is cheap and doesn’t make a copy of the
underlying data structure.
