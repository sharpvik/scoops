package bytes

import (
    "github.com/sharpvik/scoops/Package/DataTypes/Stack"
)



type Object interface {
    Print()
    Type() string
}


type (
    Instruction struct {
        opcode, operand byte
    }
    
    Environment struct {
        name string
        data *stack.Stack
        refs []*Object
        prev *Environment
        // next *Environment // ?
    }
    
    Interpreter struct {
        running bool
        err     error
        ip      uint64
        code    []*Instruction
        global  *Environment
        scope   *Environment
    }
)



func NewInstruction(opcode, operand byte) *Instruction {
    return &Instruction{opcode, operand}
}


func NewEnvironment(name string, prev *Environment) *Environment {
    return &Environment{name, stack.New(), nil, prev}
}


func NewInterpreter(code []*Instruction) *Interpreter {
    global := NewEnvironment("global", nil)
    return &Interpreter{
        true,
        nil,
        0,
        code,
        global,
        global,
    }
}