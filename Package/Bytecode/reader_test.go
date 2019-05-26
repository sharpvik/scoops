package bytecode

import (
    "os"
    "testing"
)



func TestRead(t *testing.T) {
    sampleFilename := "test.scpb"
    file, err := os.Create(sampleFilename)
    if err != nil {
        t.Error("System Error: Cannot create sample file.")
    }

    input := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    file.Write(input)

    output, err := Read(sampleFilename)

    if err != nil {
        t.Error("Fail: Cannot read sample file.")
    }

    for i, b := range input {
        if b != output[i] {
            t.Error("Fail: Cannot properly read bytecode file.")
        }
    }

    file.Close()
    os.Remove(sampleFilename)
}

