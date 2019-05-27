package bytes

import (
    "os"
    "bufio"
    "github.com/sharpvik/scoops/Package/Shared"
    //"github.com/sharpvik/scoops/Package/Util"
)



// FUNCTIONS
func Read(filename string) ([]Instruction, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    rdr := bufio.NewReader(file)
    opcode, err := rdr.ReadByte()
    if err != nil {
        return nil, err
    }
    code := []Instruction{}
    for opcode != shared.THE_END {
        //
    }
    return code, err
}
