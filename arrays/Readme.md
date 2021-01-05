# Arrays

Array is a composite data type in Go lang, lets declare an empty array

```go
// declaring an empty array of strings

var days []string

// declaring an array with elements

days := [...]string{"monday", "tuesday", "wednesday","thursday", "friday","saturday", "sunday"}
```

In order to query the second or any specific element, we can do so in a very similar way in other lang like Javascript

```go
fmt.Println(days[0]) // prints 'monday'

fmt.Println(days[5]) // prints 'saturday'
```
