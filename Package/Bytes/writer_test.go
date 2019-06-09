package bytes

import (
    "os"
    "testing"
    "io/ioutil"
)

func TestWrite(t *testing.T) {
    sampleFile := "sample.scpb"
    input := []*Instruction{
        NewInstruction(1, 1),
        NewInstruction(2, 2),
        NewInstruction(3, 3),
        NewInstruction(4, 4),
        NewInstruction(5, 5),
    }
    digest := []byte{
        1, 1,
        2, 2,
        3, 3,
        4, 4,
        5, 5,
    }
    
    err := Write(input, sampleFile)
    if err != nil {
        t.Error(err)
    }
    
    output, err := ioutil.ReadFile(sampleFile)
    if err != nil {
        t.Error(err)
    }
    if len(digest) != len(output) {
        t.Error("Cannot properly write bytecode.")
    }
    for i, b := range digest {
        if b != output[i] {
            t.Error("Cannot properly write bytecode.")
        }
    }
    
    os.Remove(sampleFile)
}