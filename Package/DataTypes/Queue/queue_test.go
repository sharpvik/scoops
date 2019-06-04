package queue

import "testing"



func TestPush(t *testing.T) {
    q := New()
    q.Push(1)
    q.Push(2)
    q.Push(3)
    if q.Size() != 3 {
        t.Error("Method Push failed test.")
    }
}


func TestEmpty(t *testing.T) {
    q := New()
    a := q.Empty()
    if !a {
        t.Error("Queue is empty but not reported as such.")
    }
    
    q.Push(1)
    b := q.Empty()
    if b {
        t.Error("Queue is not empty, but reported a empty.")
    }
}


func TestClear(t *testing.T) {
    q := New()
    q.Push(1)
    q.Push(2)
    q.Push(3.14)
    q.Clear()
    if !q.Empty() {
        t.Error("Clear function failed to clear the queue.")
    }
}


func TestPop(t *testing.T) {
    q := New()
    cases := []interface{}{
        1, true, 3.14, false, 'c',
    }
    for _, c := range cases {
        q.Push(c)
    }
    for _, j := range cases {
        i := q.Pop()
        if i != j {
            t.Error("Function Pop returned incorrect value.")
        }
    }
}


func TestPeek(t *testing.T) {
    q := New()
    cases := []interface{}{
        1, true, 3.14, false, 'c',
    }
    for _, c := range cases {
        q.Push(c)
    }
    for _, j := range cases {
        i := q.Peek()
        if i != j {
            t.Error("Function Peek returned incorrect value.")
        }
        q.Pop()
    }
}
