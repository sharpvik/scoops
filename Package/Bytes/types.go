package bytes

import (
    "bufio"
    "github.com/sharpvik/scoops/Package/Shared"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
    "github.com/sharpvik/scoops/Package/DataTypes/Stack"
    "os"
)



type (
    Instruction struct {
        opcode, operand byte
    }

    Environment struct {
        name    string
        data    *stack.Stack
        vars    []*shared.Object
        prev    *Environment
        global  *Environment
    }

    Interpreter struct {
        running bool
        err     *primitives.Error
        ip      uint64
        code    []*Instruction
        global  *Environment
        scope   *Environment    // current execution scope
        thenil  *primitives.Nil // universal nil value
        stdout  *bufio.Writer
        writer  *bufio.Writer   // current writer (stdout by default)
    }
)



func NewInstruction(opcode, operand byte) *Instruction {
    return &Instruction{opcode, operand}
}


func NewEnvironment(name string, prev, global *Environment) *Environment {
    return &Environment{name, stack.New(), nil, prev, global}
}


func NewInterpreter(code []*Instruction) *Interpreter {
    global := NewEnvironment("global", nil, nil)
    stdout := bufio.NewWriter(os.Stdout)
    return &Interpreter{
        true,
        nil,
        0,
        code,
        global,
        global,
        primitives.NewNil(),
        stdout,
        stdout,
    }
}
