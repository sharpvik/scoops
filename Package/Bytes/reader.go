package bytes

import (
    "bufio"
    "os"
)



// FUNCTIONS
func Read(filename string) ([]*Instruction, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    
    rdr := bufio.NewReader(file)
    
    var code []*Instruction
    for {
        opcode, err := rdr.ReadByte()
        if err != nil {
            return nil, err
        }
        
        operand, err := rdr.ReadByte()
        if err != nil {
            return nil, err
        }
        
        code = append( code, NewInstruction(opcode, operand) )
    }
    
    return code, nil
}
