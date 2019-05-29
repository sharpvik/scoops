package bytes

import (
    "errors"
    "fmt"
    //"github.com/sharpvik/scoops/Package/Util"
)



type Instruction struct{
    opcode  uint8
    operand []byte
}


type Interpreter struct {
    running bool
    err     error
    ip      uint64
    code    []Instruction
    // global environment
}



func NewInstruction(opcode uint8, operand []byte) *Instruction {
    return &Instruction{opcode, operand}
}


func NewInterpreter(code []Instruction) *Interpreter {
    return &Interpreter{
        true,
        nil,
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
