package bytes

import (
    "bufio"
    "os"
    "github.com/sharpvik/scoops/Package/Shared"
)



func Read(filename string) ([]*shared.Instruction, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    
    rdr := bufio.NewReader(file)
    
    var code []*shared.Instruction
    for {
        opcode, err := rdr.ReadByte()
        if err != nil {
            return nil, err
        }
        
        operand, err := rdr.ReadByte()
        if err != nil {
            return nil, err
        }
        
        code = append( code, shared.NewInstruction(opcode, operand) )
    }
    
    return code, nil
}
