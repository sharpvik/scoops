package assembly

import (
    "errors"
    "github.com/sharpvik/scoops/Package/Bytes"
    "github.com/sharpvik/scoops/Package/Shared"
    "regexp"
    "strconv"
)



func SyntaxCheck(line string) error {
    /* 
     * [A-Z_]+       -- OPERATOR
     * '[\x00-\x7F]' -- ASCII CHAR CONVERSION OPERAND
     * [0-9]+        -- DECIMAL OPERAND
     * x[0-9A-F]+    -- HEXADECIMAL OPERAND
     * b[01]{1,8}    -- BINARY OPERAND
     */
    validInstruction := regexp.MustCompile(
        `^[A-Z_]+( '[\x00-\x7F]'| [0-9]+| x[0-9A-F]+| b[01]{1,8})*\s*$`)
    if !validInstruction.MatchString(line) {
        return errors.New("Syntactically incorrect line detected.")
    }
    return nil
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


func SemanticsCheck(line string) error {
    opcodeString := FindOpcode(line)
    opcode, exists := OpcodeMap[opcodeString]
    if !exists {
        return errors.New("Opcode does not exist.")
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
        return errors.New("Operand contains byte value out of range.")
    }
    
    // Check number of bytes in operand for opcode.
    validNumberOfBytesInOperand := shared.NumberOfBytesInOperand[opcode]
    operand := FindOperand(line)
    actualNumberOfBytesInOperand := len(operand)
    if validNumberOfBytesInOperand == -1 {
        tmp, _ := strconv.ParseUint(operand[0], 10, 8)
        validNumberOfBytesInOperand = int(tmp) + 1
    }
    if actualNumberOfBytesInOperand != validNumberOfBytesInOperand {
        return errors.New("Invalid number of bytes in operand.")
    }
    
    return nil
}


func GetIntegerAndBase(integer string) (string, int) {
    /* 
     * [A-Z_]+       -- OPERATOR
     * '[\x00-\x7F]' -- ASCII CHAR CONVERSION OPERAND
     * [0-9]+        -- DECIMAL OPERAND
     * x[0-9A-F]+    -- HEXADECIMAL OPERAND
     * b[01]{1,8}    -- BINARY OPERAND
     */
    prefix := integer[0]
    switch prefix {
    case 'b':
        return integer[1:], 2
    case 'x':
        return integer[1:], 16
    default:
        return integer, 10
    }
}


func AssembleLine(line string) (instruction *bytes.Instruction, err error) {
    err = SyntaxCheck(line)
    if err != nil {
        return
    }
    err = SemanticsCheck(line)
    if err != nil {
        return
    }
    
    opcodeString := FindOpcode(line)
    operandSlice := FindOperand(line)
    opcode := OpcodeMap[opcodeString]
    var operand []byte
    for _, item := range operandSlice {
        if item[0] == '\'' { // char
            operand = append(operand, item[1])
            continue
        }
        // If this byte is not a char, it must be an integer with base
        // 2 / 10 / 16. Let's detect its base and parse it.
        integer, integerBase := GetIntegerAndBase(item)
        operand64, _ := strconv.ParseUint(integer, integerBase, 8)
        operandByte := uint8(operand64)
        operand = append(operand, operandByte)
    }
    
    instruction = bytes.NewInstruction(opcode, operand)
    return
}


/*
 * This function should work in a linear manner, checking line by line,
 * returning error whenever it finds a line that is either syntactically or
 * semantically incorrect.
 */
func Assemble(assemblyCode []string) (byteCode []*bytes.Instruction, 
                                      err error) {
    for _, line := range assemblyCode {
        instruction, err := AssembleLine(line)
        if err != nil {
            return nil, err
        }
        byteCode = append(byteCode, instruction)
    }
    return
}
