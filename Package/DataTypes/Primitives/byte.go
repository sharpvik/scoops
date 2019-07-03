package primitives

import (
    "fmt"
)



type Byte struct {
    value byte
}


func NewByte(value byte) *Byte {
    return &Byte{value}
}


func (b *Byte) Print() {
    fmt.Printf("%d", b.value)
}


func (b *Byte) Type() string {
    return "byte"
}


func AddByte(a, b *Byte) *Byte {
    return NewByte(a.value + b.value)
}


func SubByte(a, b *Byte) *Byte {
    return NewByte(a.value - b.value)
}


func MulByte(a, b *Byte) *Byte {
    return NewByte(a.value * b.value)
}


func DivByte(a, b *Byte) *Byte {
    return NewByte(a.value / b.value)
}


func (x *Byte) ToBoolean() *Boolean {
    return NewBoolean(x.value)
}


func (x *Byte) ToInteger() *Integer {
    return NewInteger( int64(x.value) )
}


func (x *Byte) ToRune() *Rune {
    return NewRune( []byte{x.value} )
}
