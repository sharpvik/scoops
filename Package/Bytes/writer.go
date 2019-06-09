package bytes

import (
    "bufio"
    "os"
)

func Write(byteCode []*Instruction, filename string) error {
    // Convert []*Instruction to []byte
    var digest []byte
    for _, i := range byteCode {
        digest = append(digest, i.opcode, i.operand)
    }
    
    // Create file
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    
    // Write to file
    wrtr := bufio.NewWriter(file)
    _, err = wrtr.Write(digest)
    if err != nil {
        return err
    }
    wrtr.Flush()
    file.Close()
    
    return nil
}
