# What is a pointer?
A pointer is a variable which stores the memory address of another variable.
Pointer variables are great for sharing variables between functions.
They allow functions to access and change the state of a variable that was declared
within the scope of a different function and possibly a different goroutine. 
## Declaring pointers

\*T is the type of the pointer variable which points to a value of type T.

## Zero value of a pointer

Zero value of a pointer is nil

## Creating pointers using the new function

Go provides a <code>new</code> function which can be used to create pointers.
The function takes a type as argument and returns a pointer to a newly allocated zero value of the type passed as argument.

## Dereferencing a pointer

Dereferencing a pointer means accessing the value of the variable which the pointer points to. \*a is the syntax to deference a.
