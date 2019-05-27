package assembly

import (
    "fmt"
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
        t.Error("Fail: Cannot properly analyse syntax.")
        fmt.Println("Got badLines:", badLines)
        fmt.Println("While realBadLines:", realBadLines)
    }
    for i, verdict := range badLines {
        if verdict != realBadLines[i] {
            t.Error("Fail: Cannot properly analyse syntax.")
            fmt.Println("Verdict for line", i, "is incorrect.")
        }
    }
}
