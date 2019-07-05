package primitives

import "fmt"



type Nil struct {
    value byte
}


func NewNil() *Nil {
    return &Nil{0}
}


func (n *Nil) Print() {
    fmt.Print("nil")
}


func (n *Nil) Type() string {
    return "nil"
}