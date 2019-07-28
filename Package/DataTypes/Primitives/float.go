package primitives

import (
    "bufio"
    "fmt"
    "github.com/sharpvik/scoops/Package/Shared"
)

// When implemented, change the README.md file in this folder.
type Float struct {
    Value float64
}


func NewFloat(value float64) *Float {
    return &Float{value}
}


func (f *Float) Clone() shared.Object {
    return NewFloat(f.Value)
}


func (f *Float) Print(w *bufio.Writer) {
    w.WriteString( fmt.Sprintf("%v", f.Value) )
}


func (f *Float) Type() string {
    return "flt"
}


func AddFloat(a, b *Float) *Float {
    return NewFloat(a.Value + b.Value)
}


func SubFloat(a, b *Float) *Float {
    return NewFloat(a.Value - b.Value)
}


func MulFloat(a, b *Float) *Float {
    return NewFloat(a.Value * b.Value)
}


func DivFloat(a, b *Float) *Float {
    return NewFloat(a.Value / b.Value)
}