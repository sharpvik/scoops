package bytes

import (
    "fmt"
)



type Integer struct {
    value int
}


func NewInteger(value int) *Integer {
    return &Integer{"int", value}
}


func (integer *Integer) Print() {
    fmt.Printf("%d", integer.value)
}


func (integer *Integer) Type() string {
    return "int"
}