package assembly

import (
    "testing"
)



func TestSyntaxCheck(t *testing.T) {
    cases := []string{
        "PUSH_CONST 'C'\n",
        "PUSH_CONST 'A' 'b'\n",
        "PUSH_CONST 1 42\n",
        "EMIT_CONST\n",
        "THE_END\n",
        
        "push_const 'c\n",
        "PUSH_CONST 'C\n",
        "PUSH_CONST 'sc'\n",
        "PUSH_CONST 'sc\n",
        "push_const 'sc'",
    }
    badLines, _ := SyntaxCheck(cases)
    realBadLines := []uint64{5, 6, 7, 8, 9}
    if len(badLines) != len(realBadLines) {
        t.Error(
            "Fail: Cannot properly analyse syntax.\n" +
            "Got badLines:", badLines, "\n" +
            "While realBadLines:", realBadLines,
        )
        return
    }
    for i, verdict := range badLines {
        if verdict != realBadLines[i] {
            t.Error(
                "Fail: Cannot properly analyse syntax.\n" +
                "Verdict for line", i, "is incorrect.",
            )
        }
    }
}


func TestFindOpcode(t *testing.T) {
    cases := []string{
        "PUSH_CONST 'C'\n",
        "PUSH_CONST 'A' 'b'\n",
        "PUSH_CONST 1 42\n",
        "EMIT_CONST\n",
        "THE_END\n",
    }
    answers := []string{
        "PUSH_CONST",
        "PUSH_CONST",
        "PUSH_CONST",
        "EMIT_CONST",
        "THE_END",
    }
    for i, c := range cases {
        reply := FindOpcode(c)
        if reply != answers[i] {
            t.Error(
                "Fail: Cannot properly find opcode.\n" +
                "Got opcode:", reply, "\n" +
                "But expected:", answers[i], 
            )
        }
    }
}


func TestFindOperand(t *testing.T) {
    cases := []string{
        "PUSH_CONST 'C'\n",
        "PUSH_CONST 'A' 'b'\n",
        "PUSH_CONST 1 42\n",
        "EMIT_CONST\n",
        "THE_END\n",
    }
    answers := [][]string{
        []string{"'C'"},
        []string{"'A'", "'b'"},
        []string{"1", "42"},
        []string{},
        []string{},
    }
    for i, c := range cases {
        reply := FindOperand(c)
        if len(reply) != len(answers[i]) {
            t.Error(
                "Fail: Cannot properly find operand.\n" +
                "Case:", i, "\n" +
                "Got operand:", reply, "\n" +
                "But expected:", answers[i], 
            )
            return
        }
        for j, item := range reply {
            if item != answers[i][j] {
                t.Error(
                    "Fail: Cannot properly find operand.\n" +
                    "Case:", i, "\n" +
                    "Got opcode:", reply, "\n" +
                    "But expected:", answers[i], 
                )
            }
        }
    }
}


func TestFindIntegers(t *testing.T) {
    cases := []string{
        "PUSH_CONST 'C'\n",
        "PUSH_CONST 5 'b'\n",
        "PUSH_CONST 1 42\n",
        "EMIT_CONST\n",
        "THE_END\n",
    }
    answers := [][]string{
        []string{},
        []string{"5"},
        []string{"1", "42"},
        []string{},
        []string{},
    }
    for i, c := range cases {
        reply := FindIntegers(c)
        if len(reply) != len(answers[i]) {
            t.Error(
                "Fail: Cannot properly find integer.\n" +
                "Case:", i, "\n" +
                "Got operand:", reply, "\n" +
                "But expected:", answers[i], 
            )
            return
        }
        for j, item := range reply {
            if item != answers[i][j] {
                t.Error(
                    "Fail: Cannot properly find integer.\n" +
                    "Case:", i, "\n" +
                    "Got opcode:", reply, "\n" +
                    "But expected:", answers[i], 
                )
            }
        }
    }
}


func TestSemanticsCheck(t *testing.T) {
    cases := []string{
        "PUSH_CONST 1 'C'\n",
        "PUSH_CONST 2 'A' 255\n",
        "PUSH_CONST 1 42\n",
        "EMIT_CONST\n",
        "THE_END\n",
        
        "END\n", // opcode does not exist
        "PUSH_CONST 256\n", // 256 cannot be stored in a uint8
        "BINARY_OP '&' '&' '&'\n", // BINARY_OP requres 2 bytes in operand not 3
        "PUSH_CONST 2 1\n", // PUSH_CONST wants 2 + 1 bytes in operand not 2
    }
    badLines, _ := SemanticsCheck(cases)
    realBadLines := []uint64{5, 6, 7, 8}
    if len(badLines) != len(realBadLines) {
        t.Error(
            "Fail: Cannot properly analyse semantics.\n" +
            "Got badLines:", badLines, "\n" +
            "While realBadLines:", realBadLines,
        )
        return
    }
    for i, line := range badLines {
        if line != realBadLines[i] {
            t.Error(
                "Fail: Cannot properly analyse semantics.\n" +
                "Got badLines:", badLines, "\n" +
                "While realBadLines:", realBadLines,
            )
        }
    }
}


func TestAssemble(t *testing.T) {
    //
}
