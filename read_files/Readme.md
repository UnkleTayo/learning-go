# Reading Files

In order to read from files on your local filesystem, you’ll have to use the io/ioutil module. You’ll first have to pull of the contents of a file into memory by calling ioutil.ReadFile("/path/to/my/file.ext") which will take in the path to the file you wish to read in as it’s only parameter. This will return either the data of the file, or an err which can be handled as you normally handle errors in go.

Create a new file called main.go as well as another file called localfile.data. Add a random piece of text to the .data file so that our finished go program has something to read and then do the following:

# Writing Files to New Files

Now that we’ve covered reading from files in Go, it’s time to look at creating and writing to our own files!

In order to write content to files using Go, we’ll again have to leverage the io/ioutil module. We’ll first have to construct a byte array that represents the content we wish to store within our files.

```Go
mydata := []byte("all my data I want to write to a file")
```

Once we have constructed this byte array, we can then call ioutil.WriteFile() to write this byte array to a file. The WriteFile() method takes in 3 different parameters, the first is the location of the file we wish to write to, the second is our mydata object, and the third is the FileMode, which represents our file’s mode and permission bits.
