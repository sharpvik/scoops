package bytes

import (
    "fmt"
)



type Integer struct {
    _type   string
    value   int
}


func NewInteger(value int) *Integer {
    return &Integer{"int", value}
}


func (integer *Integer) Print() {
    fmt.Printf("%d", integer.value)
}