# Functions

A function can be described as a block of code that performs swcific task.

## Function Declaration

The function declaration starts with a func keyword followed by the functionname. The parameters are specified between ( and ) followed by the returntype of the function.

## Variadic function

variadic function is a function that accepts a variable number of arguments. If the last parameter of a function definition is prefixed by ellipsis ..., then the function can accept any number of arguments for that parameter.

_syntax_

```Go
func hello(a int, b ...int) {
}
```

## Append is a variadic function

Append function in the standard library is a variadic function because it accepts any number of argument
