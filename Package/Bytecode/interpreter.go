package bytecode

import (
    "fmt"
    //"github.com/sharpvik/scoops/Package/Util"
)



type interpreter struct {
    running bool
    err     bool
    ip      uint64
    code    []uint8
    // global environment
}


func NewInterpreter(code []uint8) *interpreter {
    return &interpreter{
        true,
        false,
        0,
        code,
    }
}


func Run(code []uint8) {
    interpreter := NewInterpreter(code)
    for interpreter.running && !interpreter.err {
        fmt.Println("It's all good.")
        interpreter.err = true
    }
}

