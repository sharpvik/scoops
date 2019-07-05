package primitives

import (
    "fmt"
)



type Boolean struct {
    Value bool
}


func NewBoolean(value byte) *Boolean {
    return &Boolean{value != byte(0)}
}


func (b *Boolean) Print() {
    fmt.Printf("%v", b.Value)
}


func (b *Boolean) Type() string {
    return "bln"
}


func NotBoolean(b *Boolean) *Boolean {
    return &Boolean{!b.Value}
}


func AndBoolean(a, b *Boolean) *Boolean {
    return &Boolean{a.Value && b.Value}
}


func OrBoolean(a, b *Boolean) *Boolean {
    return &Boolean{a.Value || b.Value}
}


func XorBoolean(a, b *Boolean) *Boolean {
    return &Boolean{a.Value != b.Value}
}
