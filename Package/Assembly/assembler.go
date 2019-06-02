package assembly

import (
    "errors"
    "github.com/sharpvik/scoops/Package/Shared"
    "github.com/sharpvik/scoops/Package/Bytes"
    "regexp"
    "strconv"
)



func SyntaxCheck(code []string) (badLines []uint64, err error) {
    /* 
     * [A-Z_]+       -- OPERATOR
     * '[\x00-\x7F]' -- ASCII CHAR CONVERSION OPERAND
     * [0-9]+        -- DECIMAL OPERAND
     * x[0-9A-F]+    -- HEXADECIMAL OPERAND
     * b[01]{1,8}    -- BINARY OPERAND
     */
    validInstruction := regexp.MustCompile(
        `^[A-Z_]+( '[\x00-\x7F]'| [0-9]+| x[0-9A-F]+| b[01]{1,8})*\s*$`)
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
    validOperandByte := regexp.MustCompile(
        `'[\x00-\x7F]'|[0-9]+|x[0-9A-F]+|b[01]{1,8}`)
    return validOperandByte.FindAllString(line, -1)
}


func FindDecimals(line string) []string {
    validInteger := regexp.MustCompile(`[0-9]+`)
    return validInteger.FindAllString(line, -1)
}


func FindHexadecimals(line string) []string {
    validHexadecimal := regexp.MustCompile(`x[0-9A-F]+`)
    return validHexadecimal.FindAllString(line, -1)
}


func DecimalCheck(decimal string) bool {
    _, err := strconv.ParseUint(decimal, 10, 8)
    return err == nil
}


func DecimalSliceCheck(decimals []string) bool {
    for _, integer := range decimals {
        if !DecimalCheck(integer) {
            return false
        }
    }
    return true
}


func HexadecimalCheck(hexadecimal string) bool {
    _, err := strconv.ParseUint(hexadecimal[1:], 16, 8)
    return err == nil
}


func HexadecimalSliceCheck(hexadecimals []string) bool {
    for _, integer := range hexadecimals {
        if !HexadecimalCheck(integer) {
            return false
        }
    }
    return true
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
        
        // Check that operand is made of values that can be stored in a byte
        /* 
         * ASCII characters are by definition a byte long.
         * Binary integers' length is checked syntactically as well.
         * Must only check:
         *     1. Decimals
         *     2. Hexadecimals
         */
        decimalStrings := FindDecimals(line)
        hexadecimalStrings := FindHexadecimals(line)
        
        if !DecimalSliceCheck(decimalStrings) ||
           !HexadecimalSliceCheck(hexadecimalStrings) {
            errSlice = append(
                errSlice,
                errors.New("Operand contains byte value out of range."),
            )
            badLines = append( badLines, uint64(i) )
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
