package bytes

import (
    "bufio"
    "errors"
    "os"
)

func Write(assemblyCode []*Instruction, filename string) error {
    // Convert []*bytes.Instruction to []byte
    var byteCode []byte
    for _, i := range assemblyCode {
        byteCode = append(byteCode, i.Opcode)
        for _, b := range i.Operand {
            byteCode = append(byteCode, b)
        }
    }
    
    // Create file
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    
    // Create bufio.Writer and write to file
    wrtr := bufio.NewWriter(file)
    n, err := wrtr.Write(byteCode)
    if err != nil {
        return err
    }
    
    return nil
}
