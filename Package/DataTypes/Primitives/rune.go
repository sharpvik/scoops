package primitives

import (
    "fmt"
)



type Rune struct {
    value []byte
}


func NewRune(value []byte) *Rune {
    return &Rune{value}
}


func (r *Rune) Print() {
    fmt.Print( string(r.value) )
}


func (r *Rune) Type() string {
    return "rune"
}
