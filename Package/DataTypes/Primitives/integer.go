package primitives

import (
    "fmt"
)



type Integer struct {
    value int64
}


func NewInteger(value int64) *Integer {
    return &Integer{value}
}


func (integer *Integer) Print() {
    fmt.Printf("%d", integer.value)
}


func (integer *Integer) Type() string {
    return "int"
}


func AddInteger(a, b *Integer) *Integer {
    return NewInteger(a.value + b.value)
}


func SubInteger(a, b *Integer) *Integer {
    return NewInteger(a.value - b.value)
}


func MulInteger(a, b *Integer) *Integer {
    return NewInteger(a.value * b.value)
}
