package primitives

import "fmt"



type Byte struct {
    Value byte
}


func NewByte(value byte) *Byte {
    return &Byte{value}
}


func (b *Byte) Print() {
    fmt.Printf("%d", b.Value)
}


func (b *Byte) Type() string {
    return "byte"
}


func AddByte(a, b *Byte) *Byte {
    return NewByte(a.Value + b.Value)
}


func SubByte(a, b *Byte) *Byte {
    return NewByte(a.Value - b.Value)
}


func MulByte(a, b *Byte) *Byte {
    return NewByte(a.Value * b.Value)
}


func DivByte(a, b *Byte) *Byte {
    return NewByte(a.Value / b.Value)
}


func (x *Byte) ToBoolean() *Boolean {
    return NewBoolean(x.Value)
}


func (x *Byte) ToInteger() *Integer {
    return NewInteger( int64(x.Value) )
}


func (x *Byte) ToRune() *Rune {
    return NewRune( []byte{x.Value} )
}
