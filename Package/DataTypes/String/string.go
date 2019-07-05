package _string

import "github.com/sharpvik/scoops/Package/DataTypes/Primitives"



type String struct {
    size    uint64
    value   []*primitives.Rune
}


func New(value []*primitives.Rune) *String {
    return &String{uint64( len(value) ), value}
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
