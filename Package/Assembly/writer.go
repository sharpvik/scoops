package assembly

import (
    "bufio"
    "errors"
    "github.com/sharpvik/scoops/Package/Bytes"
    "os"
)

func Write(assemblyCode []*bytes.Instruction, filename string) error {
    // Convert []*bytes.Instruction to []byte
    var byteCode []byte
    for _, i := range assemblyCode {
        byteCode = append(byteCode, i.Opcode)
        for _, b := range i.Operand {
            byteCode = append(byteCode, b)
        }
    }
    
    // Write to file
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    
    wrtr := bufio.NewWriter(file)
    n, err := wrtr.Write(byteCode)
    if err != nil {
        return err
    }
    
    return nil
}
