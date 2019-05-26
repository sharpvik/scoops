package assembly

import (
    "os"
    "testing"
)



func Join(arr []string) string {
    var output string
    for _, line := range arr {
        output += line
    }
    return output
}



func TestRead(t *testing.T) {
    sampleFilename := "test.scpb"
    file, err := os.Create(sampleFilename)
    if err != nil {
        t.Error("System Error: Cannot create sample file.")
    }

    cases := []string{
        "PUSH_CONST 'c\n",
        "EMIT_CONST\n",
        "PRINT_NEWLINE\n",
        "THE_END",
    }
    
    file.WriteString( Join(cases) )
    
    answer, err := Read(sampleFilename)

    if err != nil {
        t.Error("Fail: Cannot read sample file.")
    }

    for i, line := range answer {
        if line != cases[i] {
            t.Error("Fail: Cannot properly read sample file.")
        }
    }

    file.Close()
    os.Remove(sampleFilename)
}