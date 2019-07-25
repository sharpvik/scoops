package assembly

import (
    "github.com/sharpvik/scoops/Package/Bytes"
    "testing"
    //"fmt"
)


func TestGetIntegerAndBase(t *testing.T) {
    cases := []string{
        "b101",
        "xFF",
        "42",
    }
    answers := []int{
        2, 16, 10,
    }
    for i, c := range cases {
        _, base := GetIntegerAndBase(c)
        if base != answers[i] {
            t.Errorf("Invalid base %d for %s.", base, c)
        }
    }
}


func TestAssembleLine(t *testing.T) {
    cases := []string{
        "PUSH_CONST b101010",
        "PUSH_CONST x2A",
        "PUSH_CONST xFF",
        "PUSH_CONST 42",
        "PUSH_CONST PUSH_CONST",
        "END 0",
    }
    answers := []*bytes.Instruction{
        bytes.NewInstruction(1, 42),
        bytes.NewInstruction(1, 42),
        bytes.NewInstruction(1, 255),
        bytes.NewInstruction(1, 42),
        bytes.NewInstruction(1, 1),
        bytes.NewInstruction(0, 0),
    }
    for i, c := range cases {
        instruction, err := AssembleLine(c)
        if err != nil {
            t.Error(err)
        }
        if *instruction != *answers[i] {
            t.Errorf("Cannot properly assemble line %d.", i + 1)
        }
    }
    cases = []string{
        "push_const 1",
        "PUSH_CONST 11 42",
        "GO_BROKE",
    }
    for i, c := range cases {
        _, err := AssembleLine(c)
        if err == nil {
            t.Errorf("Cannot properly detect errors (line %d).", i + 1)
        }
    }
}