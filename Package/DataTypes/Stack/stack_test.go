package stack

import "testing"



func TestPush(t *testing.T) {
    s := New()
    s.Push(true)
    s.Push(2)
    s.Push(3.14)
    if s.size != 3 {
        t.Error("Function 'Push' works incorrectly.")
    }
}


func TestEmpty(t *testing.T) {
    s := New()
    if !s.Empty() {
        t.Error("Function 'Empty' works incorrectly.")
    }
    s.Push(true)
    if s.Empty() {
        t.Error("Function 'Empty' works incorrectly.")
    }
}


func TestClear(t *testing.T) {
    s := New()
    s.Push(true)
    s.Push(2)
    s.Push(3.14)
    s.Clear()
    if !s.Empty() {
        t.Error("Funciton 'Clear' works incorrectly.")
    }
}


func TestPop(t *testing.T) {
    s := New()
    s.Push(true)
    s.Push(2)
    s.Push(3.14)
    v := s.Pop()
    if s.Size() != 2 || v != 3.14 {
        t.Error("Function 'Pop' works incorrectly.")
    }
    s.Pop()
    v = s.Pop()
    if !s.Empty() || v != true {
        t.Error("Function 'Pop' works incorrectly.")
    }
}