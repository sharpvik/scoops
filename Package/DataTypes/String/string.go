package _string

import (
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
    "github.com/sharpvik/scoops/Package/Shared"
)



type String struct {
    size    uint64
    value   []*primitives.Rune
}


func New(value []*primitives.Rune) *String {
    return &String{uint64( len(value) ), value}
}


func FromString(from string) *String {
    var value []*primitives.Rune
    runes := []rune(from)
    for _, r := range runes {
        value = append(   value, primitives.NewRune(  []byte( string(r) )  )   )
    }
    return New(value)
}


func (s *String) Clone() shared.Object {
    var value []*primitives.Rune
    for _, r := range s.value {
        value = append( value, r.Clone().(*primitives.Rune) )
    }
    return New(value)
}


func (s *String) Print() {
    for _, r := range s.value {
        r.Print()
    }
}


func (s *String) Type() string {
    return "string"
}


func (s *String) Size() uint64 {
    return s.size
}


func Concatenate(a, b *String) *String {
    value := a.value
    for _, r := range b.value {
        value = append(value, r)
    }
    return New(value)
}
