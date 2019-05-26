package bytecode

import (
    "os"
    "bufio"
    "github.com/sharpvik/scoops/Package/Bytecode"
    "github.com/sharpvik/scoops/Package/Shared"
    "github.com/sharpvik/scoops/Package/Util"
)



// FUNCTIONS
func Read(filename string) ([]bytecode.instruction, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    rdr := bufio.NewReader(file)
    opcode, err := rdr.ReadByte()
    if err != nil {
        return nil, err
    }
    code := []bytecode.instruction{}
    for opcode != shared.THE_END {
        //
    }
    return code, err
}

