package _string

import (
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
)



type String struct {
    value []*primitives.Rune
}


func New(value []*primitives.Rune) *String {
    return &String{value}
}


func (s *String) Print() {
    for _, r := range s.value {
        r.Print()
    }
}


func (s *String) Type() string {
    return "str"
}