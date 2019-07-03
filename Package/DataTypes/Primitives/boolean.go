package primitives

import (
    "fmt"
)



type Boolean struct {
    value bool
}


func NewBoolean(value byte) *Boolean {
    return &Boolean{value != byte(0)}
}


func (b *Boolean) Print() {
    fmt.Printf("%v", b.value)
}


func (b *Boolean) Type() string {
    return "bln"
}


func NotBoolean(b *Boolean) *Boolean {
    return NewBoolean(!b.value)
}


func AndBoolean(a, b *Boolean) *Boolean {
    return NewBoolean(a.value && b.value)
}


func OrBoolean(a, b *Boolean) *Boolean {
    return NewBoolean(a.value || b.value)
}


func XorBoolean(a, b *Boolean) *Boolean {
    return NewBoolean(a.value != b.value)
}
