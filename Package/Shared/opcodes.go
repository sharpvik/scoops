package shared

/*
 * If you were to ever edit the opcodes down here, make sure that you
 *     1. Run scoops/Tools/assembly_instructions_update.py
 *     2. Update...
 *         │ scoops/Package/Assembly/assembler_test.go
 *         │ scoops/InputFiles/scpa/*
 */

const (                             // │ DEC │ HEX │
    THE_END         byte = iota     // │ 0   │ x00 │
    NOP             byte = iota     // │ 1   │ x01 │
    PUSH_CONST      byte = iota     // │ 2   │ x02 │
    MAKE_BLN        byte = iota     // │ 3   │ x03 │
    MAKE_BYTE       byte = iota     // │ 4   │ x04 │
    MAKE_FLT        byte = iota     // │ 5   │ x05 │
    MAKE_INT        byte = iota     // │ 6   │ x06 │
    MAKE_NIL        byte = iota     // │ 7   │ x07 │
    MAKE_RUNE       byte = iota     // │ 8   │ x08 │
    MAKE_STR        byte = iota     // │ 9   │ x09 │
    MAKE_ARRAY      byte = iota     // │ 10  │ x0A │
    MAKE_SLICE      byte = iota     // │ 11  │ x0B │
    MAKE_LINKED     byte = iota     // │ 12  │ x0C │
    MAKE_SET        byte = iota     // │ 13  │ x0D │
    MAKE_MAP        byte = iota     // │ 14  │ x0E │
    MAKE_STACK      byte = iota     // │ 15  │ x0F │
    MAKE_QUEUE      byte = iota     // │ 16  │ x10 │
    MAKE_DEQUE      byte = iota     // │ 17  │ x11 │    
    MAKE_SCOOP      byte = iota     // │ 18  │ x12 │
    STORE_VAR       byte = iota     // │ 19  │ x13 │
    LOAD_VAR        byte = iota     // │ 20  │ x14 │
    UNARY_NOT       byte = iota     // │ 21  │ x15 │
    BINARY_AND      byte = iota     // │ 22  │ x16 │
    BINARY_OR       byte = iota     // │ 23  │ x17 │
    BINARY_XOR      byte = iota     // │ 24  │ x18 │
    BINARY_ADD      byte = iota     // │ 25  │ x19 │
    BINARY_SUB      byte = iota     // │ 26  │ x1A │
    BINARY_MUL      byte = iota     // │ 27  │ x1B │
    BINARY_DIV      byte = iota     // │ 28  │ x1C │
    BINARY_POW      byte = iota     // │ 29  │ x1D │
    PRINT_OBJ       byte = iota     // │ 30  │ x1E │
    PRINT_NEWLINE   byte = iota     // │ 31  │ x1F │
    ABSOLUTE_JUMP   byte = iota     // │ 32  │ x20 │
    RELATIVE_JUMP   byte = iota     // │ 33  │ x21 │
    JUMP_IF_TRUE    byte = iota     // │ 34  │ x22 │
    JUMP_IF_FALSE   byte = iota     // │ 35  │ x23 │
    CALL_SCOOP      byte = iota     // │ 36  │ x24 │
    RETURN          byte = iota     // │ 37  │ x25 │
    POP             byte = iota     // │ 38  │ x26 │
)
