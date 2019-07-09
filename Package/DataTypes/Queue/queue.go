package queue

import (
    "fmt"
    "github.com/sharpvik/scoops/Package/Shared"
)



type (
    node struct {
        prev    *node
        next    *node
        val     shared.Object
    }
    
    Queue struct {
        size    uint64
        head    *node
    }
)



func newNode(prev *node, next *node, val shared.Object) *node {
    return &node{prev, next, val}
}


func New() *Queue {
    return &Queue{0, nil}
}


func (q *Queue) Clear() {
    q.head = nil
    q.size = 0
}


func (q *Queue) Empty() bool {
    return q.size == 0
}


func (q *Queue) Peek() shared.Object {
    if q.size == 0 {
        return nil
    }
    return q.head.val
}


func (q *Queue) Pop() shared.Object {
    var n shared.Object
    if q.size == 0 {
        return nil
    } else if q.size == 1 {
        n = q.head.val
        q.head = nil
    } else {
        n = q.head.val
        q.head.prev.next = q.head.next
        q.head.next.prev = q.head.prev
        q.head = q.head.next
    }
    q.size--
    return n
}


func (q *Queue) Print() {
    if q.size == 0 {
        fmt.Println("<- { } <-")
        return
    }
    cur := q.head
    fmt.Print("<- {")
    for {
        fmt.Print(" ")
        cur.val.Print()
        fmt.Print(" ")
        cur = cur.next
        if cur == q.head {
            fmt.Print("} <-\n")
            break
        }
    }    
}


func (q *Queue) Push(item shared.Object) {
    n := newNode(nil, nil, item)
    if q.size == 0 {
        n.prev = n
        n.next = n
        q.head = n
    } else {
        n.prev = q.head.prev
        n.next = q.head
        q.head.prev.next = n
        q.head.prev = n
    }
    q.size++
}


func (q *Queue) Size() uint64 {
    return q.size
}


func (q *Queue) Type() string {
    return "queue"
}
