package primitives

import (
    "fmt"
)



type Boolean struct {
    value bool
}


func NewBoolean(value byte) *Boolean {
    if value != byte(0) {
        return &Boolean{true}
    } else {
        return &Boolean{false}
    }
}


func (b *Boolean) Print() {
    fmt.Printf("%v", b.value)
}


func (b *Boolean) Type() string {
    return "bln"
}