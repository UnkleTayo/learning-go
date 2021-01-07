# Slices

Slices in Go are just like arrays but with super powers this is because they allow you to access a subset of array's element

Slices are made up of three things, a pointer, a length, and a capacity.

```go
days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
weekdays := days[0:5]
fmt.Println(weekdays)
// This returns: [Monday Tuesday Wednesday Thursday Friday]
```
