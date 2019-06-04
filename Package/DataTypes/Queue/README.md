# Package queue

The implementation of the Queue abstract data type.
[Wikipedia article.](https://en.wikipedia.org/wiki/Queue_(abstract_data_type))



## Overview

This package provides the Queue data type and functions that are related to it.



## Index

```
type queue
    func New() *queue
    func (q *queue) Clear()
    func (q *queue) Empty() bool
    func (q *queue) Peek() interface{}
    func (q *queue) Pop() interface{}
    func (q *queue) Print()
    func (q *queue) Push(interface{})
    func (q *queue) Size() int
```


### func New

```go
func New() *queue
```

**New** returns pointer to a newly initialised queue.


### func Clear

```go
func (q *queue) Clear()
```

**Clear** simply empties the queue. It is *much* more efficient to use **Clear**
than popping every element off by hand using the **Pop** function.


### func Empty

```go
func (q *queue) Empty() bool
```

**Empty** returns *true* if queue is empty and *false* if it isn't.


### func Peek

```go
func (q *queue) Peek() interface{}
```

**Peek** returns the top item of the queue. Top element is said to be the
element that would be popped if **Pop** function was called.


### func Pop

```go
func (q *queue) Pop() interface{}
```

**Pop** pops off and returns the top item of the queue.


### func Print

```go
func (q *queue) Print()
```

**Print** prints out contents of the queue to stdout.


### func Push

```go
func (q *queue) Push(interface{})
```

**Push** inserts given item into the queue.


### func Size

```go
func (q *queue) Size() int
```

**Size** returns the number of elements in the queue.



## Examples

```go
package main

import (
    "fmt"
    "github.com/sharpvik/scoops/Package/DataTypes/Queue"
)

func main() {
    q := queue.New()    // init new queue and store pointer to it in 'q'
    q.Push(1)           // push first element into the queue
    q.Print()           // print queue to stdout
    one := q.Peek()     // store the top item of the queue in 'one'
    two := q.Pop()      // pop the top item off and store it in 'two'
    q.Print()           // print now empty queue to stdout
    b := (one == two)   // 'b' is true
    fmt.Println(b)      // print value of 'b' to stdout
}
```