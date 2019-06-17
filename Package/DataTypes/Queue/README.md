# Package Queue

The implementation of the Queue abstract data type.
[Wikipedia article.](https://en.wikipedia.org/wiki/Queue_(abstract_data_type))



## Overview

This package provides the Queue data type and functions that are related to it.



## Index

```
type Queue
    func New() *Queue
    func (q *Queue) Clear()
    func (q *Queue) Empty() bool
    func (q *Queue) Peek() interface{}
    func (q *Queue) Pop() interface{}
    func (q *Queue) Print()
    func (q *Queue) Push(interface{})
    func (q *Queue) Size() int
```


### func New

```go
func New() *Queue
```

**New** returns pointer to a newly initialised Queue.


### func Clear

```go
func (q *Queue) Clear()
```

**Clear** simply empties the queue. It is *much* more efficient to use **Clear**
than popping every element off by hand using the **Pop** function.


### func Empty

```go
func (q *Queue) Empty() bool
```

**Empty** returns `true` if queue is empty and `false` if it isn't.


### func Peek

```go
func (q *Queue) Peek() interface{}
```

**Peek** returns the top item of the queue. Top element is said to be the
element that would be popped if **Pop** function was called.


### func Pop

```go
func (q *Queue) Pop() interface{}
```

**Pop** pops off and returns the top item of the queue.


### func Print

```go
func (q *Queue) Print()
```

**Print** prints out contents of the queue to stdout.


### func Push

```go
func (q *Queue) Push(interface{})
```

**Push** inserts given item into the queue.


### func Size

```go
func (q *Queue) Size() int
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
    q := Queue.New()    // init new Queue and store pointer to it in 'q'
    q.Push(1)           // push first element into the queue
    q.Print()           // print queue to stdout
    one := q.Peek()     // store the top item of the queue in 'one'
    two := q.Pop()      // pop the top item off and store it in 'two'
    q.Print()           // print now empty queue to stdout
    b := (one == two)   // 'b' is true
    fmt.Println(b)      // print value of 'b' to stdout
}
```