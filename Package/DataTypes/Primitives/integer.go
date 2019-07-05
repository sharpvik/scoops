package primitives

import (
    "fmt"
)



type Integer struct {
    Value int64
}


func NewInteger(value int64) *Integer {
    return &Integer{value}
}


func (integer *Integer) Print() {
    fmt.Printf("%d", integer.Value)
}


func (integer *Integer) Type() string {
    return "int"
}


func AddInteger(a, b *Integer) *Integer {
    return NewInteger(a.Value + b.Value)
}


func SubInteger(a, b *Integer) *Integer {
    return NewInteger(a.Value - b.Value)
}


func MulInteger(a, b *Integer) *Integer {
    return NewInteger(a.Value * b.Value)
}
