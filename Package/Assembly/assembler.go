package assembly

import (
    "regexp"
    "errors"
    //"github.com/sharpvik/scoops/Package/Util"
    "github.com/sharpvik/scoops/Package/Shared"
)



var OpcodeMap = map[string]uint8{
    "THE_END": shared.THE_END,
    "NOP": shared.NOP,
    "PUSH_CONST": shared.PUSH_CONST,
    "EMIT_CONST": shared.EMIT_CONST,
    "BINARY_OP": shared.BINARY_OP,
    "COMPARE_OP": shared.COMPARE_OP,
    "PRINT_NEWLINE": shared.PRINT_NEWLINE,
}



func SyntaxCheck(code []string) (badLines []uint64, err error) {
    validInstruction := regexp.MustCompile(
        `^[A-Z_]+( '[\x21-\x7E]'| [0-9]+)*\s*$`)
    for i, line := range code {
        if !validInstruction.MatchString(line) {
            err = errors.New("Syntactically incorrect lines detected.")
            badLines = append( badLines, uint64(i) )
        }
    }
    return
}


func Assemble(code []string) {
    //
}

