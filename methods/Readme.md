# Method

A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type.

```Go
func (t Type) methodName(parameter list) {
}
```

```Go
Note - One thing you should note when working with methods is that, like functions, they create copies of the arguments passed into it. To avoid this, we can use pointer receivers when defining our methods
```
