package queue

import (
    "fmt"
    "testing"
)



func TestPrint(t *testing.T) {
    fmt.Println("\nTestPush and TestPrint:")
    q := New()
    q.Print()
    q.Push(1)
    q.Push(2)
    q.Push(3)
    q.Print()
}


func TestPeek(t *testing.T) {
    fmt.Println("\nTestPeek:")
    q := New()
    q.Push(1)
    q.Push(2)
    q.Push(3)
    n := q.Peek()
    fmt.Println(n)
}


func TestPop(t *testing.T) {
    fmt.Println("\nTestPop:")
    q := New()
    q.Push(1)
    q.Push(2)
    q.Push(3)
    q.Print()
    q.Pop()
    q.Print()
    q.Pop()
    q.Print()
}
