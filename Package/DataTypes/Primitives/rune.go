package primitives

import "fmt"



type Rune struct {
    Value []byte
}


func NewRune(value []byte) *Rune {
    return &Rune{value}
}


func (r *Rune) Print() {
    fmt.Print( string(r.Value) )
}


func (r *Rune) Type() string {
    return "rune"
}
