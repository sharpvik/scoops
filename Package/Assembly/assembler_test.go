package assembly

import (
    "testing"
)



func TestSyntaxCheck(t *testing.T) {
    cases := []string{
        "LOAD_BYTES 'C' xF7\n",
        "LOAD_BYTES 3 'c' 23 xFF b101\n",
        "LOAD_BYTES 1 42\n",
        "PRINT_OBJECT\n",
        "THE_END\n",
        
        "load_bytes 'c'\n",
        "LOAD_BYTES b0134\n",
        "LOAD_BYTES 'sc'\n",
        "LOAD_BYTES 'sc\n",
        "LOAD_BYTES xFH",
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
        "LOAD_BYTES 'C'\n",
        "LOAD_BYTES 'A' 'b'\n",
        "LOAD_BYTES 1 42\n",
        "PRINT_OBJECT\n",
        "THE_END\n",
    }
    answers := []string{
        "LOAD_BYTES",
        "LOAD_BYTES",
        "LOAD_BYTES",
        "PRINT_OBJECT",
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
        "LOAD_BYTES 'C'\n",
        "LOAD_BYTES 3 'c' 23 xFF b101\n",
        "LOAD_BYTES 1 42\n",
        "PRINT_OBJECT\n",
        "THE_END\n",
    }
    answers := [][]string{
        []string{"'C'"},
        []string{"3", "'c'", "23", "xFF", "b101"},
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
                    "Got operand:", reply, "\n" +
                    "But expected:", answers[i], 
                )
            }
        }
    }
}


func TestFindDecimals(t *testing.T) {
    cases := []string{
        "LOAD_BYTES 'C'\n",
        "LOAD_BYTES 5 'b'\n",
        "LOAD_BYTES 1 42\n",
        "PRINT_OBJECT\n",
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
        reply := FindDecimals(c)
        if len(reply) != len(answers[i]) {
            t.Error(
                "Fail: Cannot properly find decimal.\n" +
                "Case:", i, "\n" +
                "Got:", reply, "\n" +
                "But expected:", answers[i], 
            )
            return
        }
        for j, item := range reply {
            if item != answers[i][j] {
                t.Error(
                    "Fail: Cannot properly find decimal.\n" +
                    "Case:", i, "\n" +
                    "Got:", reply, "\n" +
                    "But expected:", answers[i], 
                )
            }
        }
    }
}


func TestFindHexadecimals(t *testing.T) {
    cases := []string{
        "LOAD_BYTES 'C'\n",
        "LOAD_BYTES 5 x54 xFF x65 101 'c'\n",
        "LOAD_BYTES 1 x42\n",
        "PRINT_OBJECT\n",
        "THE_END\n",
    }
    answers := [][]string{
        []string{},
        []string{"x54", "xFF", "x65", },
        []string{"x42"},
        []string{},
        []string{},
    }
    for i, c := range cases {
        reply := FindHexadecimals(c)
        if len(reply) != len(answers[i]) {
            t.Error(
                "Fail: Cannot properly find hexadecimal.\n" +
                "Case:", i, "\n" +
                "Got:", reply, "\n" +
                "But expected:", answers[i], 
            )
            return
        }
        for j, item := range reply {
            if item != answers[i][j] {
                t.Error(
                    "Fail: Cannot properly find hexadecimal.\n" +
                    "Case:", i, "\n" +
                    "Got:", reply, "\n" +
                    "But expected:", answers[i],
                )
            }
        }
    }
}


func TestDecimalCheck(t *testing.T) {
    cases := []string{
        "0", "5", "255", "128",
        "256", "-256", "1000", "-1000",
    }
    answers := []bool{
        true, true, true, true,
        false, false, false, false, 
    }
    for i, c := range cases {
        reply := DecimalCheck(c)
        if reply != answers[i] {
            t.Error(
                "Fail: Cannot properly check decimal.\n" +
                "Case:", c, "\n" +
                "Got:", reply, "\n" +
                "But expected:", answers[i],
            )
        }
    }
}


func TestHexadecimalCheck(t *testing.T) {
    cases := []string{
        "x00", "x05", "xFF",
        "x1FF", "xFFF",
    }
    answers := []bool{
        true, true, true,
        false, false,
    }
    for i, c := range cases {
        reply := HexadecimalCheck(c)
        if reply != answers[i] {
            t.Error(
                "Fail: Cannot properly check hexadecimal.\n" +
                "Case:", c, "\n" +
                "Got:", reply, "\n" +
                "But expected:", answers[i],
            )
        }
    }
}


func TestSemanticsCheck(t *testing.T) {
    cases := []string{
        "LOAD_BYTES 1 'C'\n",
        "LOAD_BYTES 4 'c' 23 xFF b101\n",
        "LOAD_BYTES 1 42\n",
        "PRINT_OBJECT\n",
        "THE_END\n",
        
        "END\n", // opcode does not exist
        "LOAD_BYTES 256\n", // 256 cannot be stored in a uint8
        "BINARY_OP '&' '&' '&'\n", // BINARY_OP requires only 2 bytes in operand
        "LOAD_BYTES 2 'c' 23 xFF\n", // LOAD_BYTES requires 2 + 1 bytes not 4
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
