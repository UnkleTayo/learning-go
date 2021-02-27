# Methods

Methods provide a way to add behavior to user-defined types. Methods are really functions that contain an extra parameter that’s declared between the keyword func and the function name.

A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type.

```Go
func (t Type) methodName(parameter list) {
}
```

## The nature of types

After declaring a new type, try to answer this question before declaring methods for the type. Does adding or removing something from a value of this type need to create a new value or mutate the existing one? If the answer is create a new value, then use value receivers for your methods. If the answer is mutate the value, then use pointer receivers. This also applies to how values of this type should be passed other parts of your program. It’s important to be consistent. The idea is to not focus on what the method is doing with the value, but to focus on what the nature of the value is.

## Reference types

Reference types in Go are the set of slice, map, channel, interface, and function types.
When you declare a variable from one of these types, the value that’s created is called a <code>header</code> value.
