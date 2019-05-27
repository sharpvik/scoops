package assembly

import (
    "regexp"
    "errors"
    "strconv"
    //"github.com/sharpvik/scoops/Package/Util"
    "github.com/sharpvik/scoops/Package/Shared"
    "github.com/sharpvik/scoops/Package/Bytes"
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
    /* 
        [A-Z_]+       -- OPERATOR
        '[\x00-\x7F]' -- CHAR CONVERSION OPERAND
        [0-9]+        -- NUMBER OPERAND                                         
    */
    validInstruction := regexp.MustCompile(
        `^[A-Z_]+( '[\x00-\x7F]'| [0-9]+)*\s*$`)
    for i, line := range code {
        if !validInstruction.MatchString(line) {
            err = errors.New("Syntactically incorrect lines detected.")
            badLines = append( badLines, uint64(i) )
        }
    }
    return
}


func FindOpcode(line string) string {
	validOpcode := regexp.MustCompile(`[A-Z_]+`)
	return validOpcode.FindString(line)
}


func FindOperand(line string) []string {
    validOperandByte := regexp.MustCompile(`'[\x00-\x7F]'|[0-9]+`)
    return validOperandByte.FindAllString(line, -1)
}


func FindIntegers(line string) []string {
    validInteger := regexp.MustCompile(`[0-9]+`)
    return validInteger.FindAllString(line, -1)
}


func SemanticsCheck(code []string) (badLines []uint64, errSlice []error) {
    for i, line := range code {
        // Check that opcode exists.
        opcodeString := FindOpcode(line)
        opcode, exists := OpcodeMap[opcodeString]
        if !exists {
            errSlice = append( errSlice, errors.New("Opcode does not exist.") )
            badLines = append( badLines, uint64(i) )
            continue
        }
        
        // Check that every integer i in operand can be stored in a uint8.
        integerStrings := FindIntegers(line)
        broken := false
        for _, integer := range integerStrings {
            _, err := strconv.ParseUint(integer, 10, 8)
            if err != nil {
                errSlice = append(
                    errSlice,
                    errors.New("Operand contains byte value out of range."),
                )
                badLines = append( badLines, uint64(i) )
                broken = true
                break
            }
        }
        if broken {
            continue
        }
        
        // Check number of bytes in operand for opcode.
        validNumberOfBytesInOperand := shared.NumberOfBytesInOperand[opcode]
        operand := FindOperand(line)
        actualNumberOfBytesInOperand := len(operand)
        if validNumberOfBytesInOperand == -1 {
            tmp, _ := strconv.ParseUint(
                operand[0], 10, 8)
            validNumberOfBytesInOperand = int(tmp) + 1
        }
        if actualNumberOfBytesInOperand != validNumberOfBytesInOperand {
            errSlice = append(
                errSlice,
                errors.New("Invalid number of bytes in operand."),
            )
            badLines = append( badLines, uint64(i) )
            continue
        }
    }
    return
}


func Assemble(code []string) (byteCode []bytes.Instruction) {
    return
}
