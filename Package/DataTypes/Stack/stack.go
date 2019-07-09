package stack

import (
    "fmt"
    "github.com/sharpvik/scoops/Package/Shared"
)



type (
    node struct {
        next    *node
        val     shared.Object
    }
    
    Stack struct {
        size    uint64
        top     *node
    }
)



func newNode(next *node, val shared.Object) *node {
    return &node{next, val}
}


func New() *Stack {
    return &Stack{0, nil}
}


func (s *Stack) Clear() {
    s.size = 0
    s.top = nil
}


func (s *Stack) Empty() bool {
    return s.size == 0
}


func (s *Stack) Peek() shared.Object {
    return s.top.val
}


func (s *Stack) Pop() shared.Object {
    if s.size == 0 {
        return nil
    }
    n := s.top.val
    s.top = s.top.next
    s.size--
    return n
}


func (s *Stack) Print() {
    if s.size == 0 {
        fmt.Println("<-> { }")
        return
    }
    cur := s.top
    fmt.Print("<-> {")
    for {
        cur.val.Print()
        cur = cur.next
        if cur == nil {
            fmt.Print("}\n")
            break
        }
    }
}


func (s *Stack) Push(item shared.Object) {
    if s.size == 0 {
        s.top = newNode(nil, item)
    } else {
        s.top = newNode(s.top, item)
    }
    s.size++
}


func (s *Stack) Size() uint64 {
    return s.size
}


func (s *Stack) Type() string {
    return "stack"
}
