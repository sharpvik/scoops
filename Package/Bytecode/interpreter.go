package bytecode

import (
    "fmt"
    "errors"
    //"github.com/sharpvik/scoops/Package/Util"
)



type Interpreter struct {
    running bool
    err     error
    ip      uint64
    code    []uint8
    // global environment
}


func NewInterpreter(code []Instruction) *interpreter {
    return &interpreter{
        true,
        false,
        0,
        code,
    }
}


func Execute(code []Instruction) {
    interpreter := NewInterpreter(code)
    for interpreter.running && interpreter.err == nil {
        fmt.Println("It's all good.")
        interpreter.err = errors.New("Just to show off.")
    }
}
