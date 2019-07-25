# Package stack

The implementation of the Stack abstract data type.
[Wikipedia article.](https://en.wikipedia.org/wiki/Stack_(abstract_data_type))



## Overview

This package provides the Stack data type and functions that are related to it.



## Index

```go
type Stack
    func New() *Stack
    func (q *Stack) Clear()
    func (s *Stack) Clone() shared.Object
    func (q *Stack) Empty() bool
    func (q *Stack) Peek() interface{}
    func (q *Stack) Pop() interface{}
    func (q *Stack) Print()
    func (q *Stack) Push(interface{})
    func (q *Stack) Size() int
    func (s *Stack) Type() string
```


### func New

```go
func New() *Stack
```

**New** returns pointer to a newly initialised Stack.


### func Clear

```go
func (q *Stack) Clear()
```

**Clear** simply empties the stack. It is *much* more efficient to use **Clear**
than popping every element off by hand using the **Pop** function.


### func Empty

```go
func (q *Stack) Empty() bool
```

**Empty** returns `true` if stack is empty and `false` if it isn't.


### func Peek

```go
func (q *Stack) Peek() interface{}
```

**Peek** returns the top item of the stack. Top element is said to be the
element that would be popped if **Pop** function was called.


### func Pop

```go
func (q *Stack) Pop() interface{}
```

**Pop** pops off and returns the top item of the stack.


### func Print

```go
func (q *Stack) Print()
```

**Print** prints out contents of the stack to stdout.


### func Push

```go
func (q *Stack) Push(interface{})
```

**Push** inserts given item into the stack.


### func Size

```go
func (q *Stack) Size() int
```

**Size** returns the number of elements in the stack.


### func Type

```go
func (s *Stack) Type() string
```

**Type** returns `Stack`'s type as Go's embedded `string` type.



## Examples

```go
package main

import (
    "fmt"
    "github.com/sharpvik/scoops/Package/DataTypes/Stack"
)

func main() {
    q := Stack.New()    // init new Stack and store pointer to it in 'q'
    q.Push(1)           // push first element into the stack
    q.Print()           // print stack to stdout
    one := q.Peek()     // store the top item of the stack in 'one'
    two := q.Pop()      // pop the top item off and store it in 'two'
    q.Print()           // print now empty stack to stdout
    b := (one == two)   // 'b' is true
    fmt.Println(b)      // print value of 'b' to stdout
}
```