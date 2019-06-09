package stack
/*
import (
    "fmt"
)
*/


type (
    node struct {
        prev    *node
        next    *node
        val     interface{}
    }
    
    Stack struct {
        size    int
        bottom  *node
    }
)



func New() *Stack {
    return &Stack{0, nil}
}
