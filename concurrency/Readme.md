# What are channels

Channels can be thought as pipes with which Goroutines communicate. Similar to how water flows from one end to another in a pipe, data can be sent from one end and received from the another end using channels

## Declaring channels

channel is defined using the make function just like slices and maps

```Go
b := make(chan int)
```

The above code creates a channel b of type int

## Sending and receiving from channel

The syntax to send and receive data from a channel are given below,

```Go
data := <- a // read from channel a
a <- data // write to channel a
```

The arrow directoion with respect to the channel shoes whether data is sent or received

In the first line, the arrow points outwards from a and hence we are reading from channel a and storing the value to the variable data.

In the second line, the arrow points towards a and hence we are writing to channel a

## Buffer Channels

Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer.

```Go
b := make(chan int, capacity)
```

For a channel to have buffer capacity should be greater than 0

## Length vs Capacity

The capacity of a buffered channel is the number of values that the channel can hold. This is the value we specify when creating the buffered channel using the make function.

The length of the buffered channel is the number of elements currently queued in it.

## WaitGroup

A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is blocked until all Goroutines finish executing. Lets say we have 3 concurrently executing Goroutines spawned from the main Goroutine. The main Goroutines needs to wait for the 3 other Goroutines to finish before terminating. This can be accomplished using WaitGroup.
