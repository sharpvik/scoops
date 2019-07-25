package primitives

import (
    "fmt"
    "github.com/sharpvik/scoops/Package/Shared"
)



type Nil struct {
    value byte
}


func NewNil() *Nil {
    return &Nil{0}
}


func (n *Nil) Clone() shared.Object {
    return n
}


func (n *Nil) Print() {
    fmt.Print("nil")
}


func (n *Nil) Type() string {
    return "nil"
}