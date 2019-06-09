package bytes



type Instruction struct{
    opcode, operand byte
}


func NewInstruction(opcode, operand byte) *Instruction {
    return &Instruction{opcode, operand}
}



type Interpreter struct {
    running bool
    err     error
    ip      uint64
    code    []*Instruction
    // global environment
}


func NewInterpreter(code []*Instruction) *Interpreter {
    return &Interpreter{
        true,
        nil,
        0,
        code,
    }
}