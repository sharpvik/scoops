package bytecode

type Instruction struct {
    opcode  uint8
    operand []byte
}

func NewInstruction(opcode uint8, operand []byte) *Instruction {
    return &Instruction{opcode, operand}
}
