package queue

import "fmt"

/*
 * Empty -> push Item
 * q.head = Item
 * q.len++
 *
 * Not empty -> push Item
 * q.head.prev.setnext(n)
 * q.head.setprev(n)
 *
 * Len == 1 -> pop Item
 * q.head = nil
 * q.len--
 *
 * Len > 1 -> pop Item
 * q.head.prev.prev.next = q.head
 * q.head.prev = q.head.prev.prev
 *
 */



type (
    Node struct {
        prev    *Node
        next    *Node
        val     interface{}
    }
    Queue struct {
        len     int
        head    *Node
    }
)



func New() *Queue {
    return &Queue{0, nil}
}


func newNode(prev *Node, next *Node, val interface{}) *Node {
    return &Node{prev, next, val}
}


func (q *Queue) Push(item interface{}) {
    n := newNode(nil, nil, item)
    if q.len == 0 {
        n.prev = n
        n.next = n
        q.head = n
    } else {
        n.prev = q.head.prev
        n.next = q.head
        q.head.prev.next = n
        q.head.prev = n
    }
    q.len++
}


func (q *Queue) Peek() interface{} {
    if q.len == 0 {
        return nil
    }
    return q.head.val
}


func (q *Queue) Pop() interface{} {
    var n interface{}
    if q.len < 1 {
        return nil
    } else if q.len == 1 {
        n = q.head.val
        q.head = nil
    } else {
        n = q.head.val
        q.head.prev.next = q.head.next
        q.head.next.prev = q.head.prev
        q.head = q.head.next
    }
    q.len--
    return n
}


func (q *Queue) Print() {
    if q.len == 0 {
        fmt.Println("{ }")
        return
    }
    cur := q.head
    fmt.Printf("{")
    for {
        fmt.Printf(" %v ", *cur)
        cur = cur.next
        if cur == q.head {
            fmt.Printf("}\n")
            break
        }
    }    
}
