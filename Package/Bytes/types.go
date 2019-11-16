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
    data    *stack.Stack
    vars    []shared.Object
    prev    *Environment
}

type Interpreter struct {
    running bool
    err     *primitives.Error
    consts  []shared.Object
    global  *Environment
    scope   *Environment    // current execution scope
    thenil  *primitives.Nil // universal nil value
    stdout  *bufio.Writer
    writer  *bufio.Writer   // current writer (stdout by default)
}



func NewEnvironment(name string, code []*shared.Instruction,
                    prev *Environment) *Environment {
    return &Environment{name, 0, code, stack.New(), nil, prev}
}


func NewInlineEnvironment(name string, ip uint64, code []*shared.Instruction,
                          prev *Environment) *Environment {
    return &Environment{name, ip, code, stack.New(), nil, prev}
}


func NewInterpreter(code []*shared.Instruction) *Interpreter {
    global := NewEnvironment("global", code, nil)
    stdout := bufio.NewWriter(os.Stdout)
    return &Interpreter{
        true,
        nil,
        nil,
        global,
        global,
        primitives.NewNil(),
        stdout,
        stdout,
    }
}
