package bytes

import (
    "os"
    "testing"
    "io/ioutil"
)

func TestWrite(t *testing.T) {
    sampleFile := "sample.scpb"
    input := []*Instruction{
        NewInstruction(1, []byte{1, 2, 3}),
        NewInstruction(2, []byte{1, 2, 3}),
        NewInstruction(3, []byte{1, 2, 3}),
        NewInstruction(4, []byte{1, 2, 3}),
        NewInstruction(5, []byte{1, 2, 3}),
    }
    digest := []byte{
        1, 1, 2, 3,
        2, 1, 2, 3,
        3, 1, 2, 3,
        4, 1, 2, 3,
        5, 1, 2, 3,
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