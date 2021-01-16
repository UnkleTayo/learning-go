# Creating a Go Module

A Go module is created byr running

```Go
go mod init packagename // most times the packgae na is the containig folders name
```

The above command will create a file called go.mod in the folder

## Init function

Each package in Go can contain an init function. The init function must not have any return type and it must not have any parameters.
The init function cannot be called explicitly in our source code. It will be called automatically when the package is initialized. The init function has the following syntax
