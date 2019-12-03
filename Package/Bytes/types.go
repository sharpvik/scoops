package bytes

import (
    "bufio"
    "github.com/sharpvik/scoops/Package/Shared"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
    "github.com/sharpvik/scoops/Package/DataTypes/Stack"
    "os"
)



type Environment struct {
    name    string
    ip      uint64          // instruction pointer
    code    []*shared.Instruction
    global  *Environment
    data    *stack.Stack
    vars    []shared.Object
    prev    *Environment
    writer  *bufio.Writer   // current writer (stdout by default)
}

type Interpreter struct {
    running bool
    err     *primitives.Error
    consts  []shared.Object
    global  *Environment
    scope   *Environment    // current execution scope
    thenil  *primitives.Nil // universal nil value
    stdout  *bufio.Writer
}



func NewEnvironment(name string, code []*shared.Instruction,
                    global, prev *Environment,
                    writer *bufio.Writer) *Environment {
    return &Environment{
        name,
        0,
        code,
        global,
        stack.New(),
        nil,
        prev,
        writer,
    }
}


func NewInlineEnvironment(name string, ip uint64, code []*shared.Instruction,
                          global, prev *Environment) *Environment {
    return &Environment{
        name,
        ip,
        code,
        global,
        stack.New(),
        nil,
        prev,
        prev.writer,
    }
}


func NewInterpreter(code []*shared.Instruction) *Interpreter {
    stdout := bufio.NewWriter(os.Stdout)
    global := NewEnvironment("global", code, nil, nil, stdout)
    return &Interpreter{
        true,
        nil,
        nil,
        global,
        global,
        primitives.NewNil(),
        stdout,
    }
}
