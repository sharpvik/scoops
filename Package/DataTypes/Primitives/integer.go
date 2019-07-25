package primitives

import (
    "fmt"
    "github.com/sharpvik/scoops/Package/Shared"
)


type Integer struct {
    Value int64
}


func NewInteger(value int64) *Integer {
    return &Integer{value}
}


func (i *Integer) Clone() shared.Object {
    return NewInteger(i.Value)
}


func (i *Integer) Print() {
    fmt.Printf("%d", i.Value)
}


func (i *Integer) Type() string {
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
