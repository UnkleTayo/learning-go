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

## Map of struct

So far we have only been storing the salary of the employee in the map. Wouldn't it be nice if we are able to store the country of each employee in the map too? This can be achieved by using a map of structs. The employee can be represented as a struct containing fields salary and country and they will be stored in the map with a string key and struct value. Let's write a program to understand how this can be done.

## Length of the map

Length of the map can be determined using the len function.
