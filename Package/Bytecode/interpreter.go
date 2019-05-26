package bytecode

import (
    "fmt"
    "errors"
    //"github.com/sharpvik/scoops/Package/Util"
)



type interpreter struct {
    running bool
    err     error
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


func Execute(code []uint8) {
    interpreter := NewInterpreter(code)
    for interpreter.running && interpreter.err == nil {
        fmt.Println("It's all good.")
        interpreter.err = errors.New("Just to show off.")
    }
}

