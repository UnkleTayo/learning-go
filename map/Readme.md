# Map

A map is a builtin type in Go that is used to store key-value pairs.
maps are just like Go's representation of hash tables data structure, it allows you to map one arbitrary type to another.

## How to create a Map

A map can be created by passing the type of key and value to the make function. The following is the syntax to create a new map.

```Go
make(map[type of key]type of value)
```

E.g

```Go
playerSalary := make(map[string]int)
```

## Deleting items from a map

delete(map, key) is the syntax to delete key from a map. The delete function does not return any value.

```Go
delete(map, key)
```

## Length of the map

Length of the map can be determined using the len function.

## Checking if a key exists

To check if an emplaoyee exist in a map the syntax is usedd

```Go
 value, ok := map[key]
```

The above is the syntax to find out whether a particular key is present in a map. If ok is true, then the key is present and its value is present in the variable value, else the key is absent.

## Maps are reference types

Similar to slices, maps are reference types. When a map is assigned to a new variable, they both point to the same internal data structure. Hence changes made in one will reflect in the other.

## Maps Equality

Maps can't be compared using the == operator. The == can be only used to check if a map is nil.

One way to check whether two maps are equal is to compare each one's individual elements one by one.
