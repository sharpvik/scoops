package shared

/*
 * If you were to ever edit the opcodes down here, make sure that you
 *     1. Run scoops/Tools/assembly_instructions_update.py
 *     2. Update...
 *         │ scoops/Package/Assembly/assembler_test.go
 *         │ scoops/InputFiles/scpa/*
 */

const (
    END             byte = iota
    BEGIN           byte = iota
    NOP             byte = iota
    PUSH_CONST      byte = iota
    MAKE_BLN        byte = iota
    MAKE_BYTE       byte = iota
    MAKE_FLT        byte = iota
    MAKE_INT        byte = iota
    MAKE_NIL        byte = iota
    MAKE_RUNE       byte = iota
    MAKE_STR        byte = iota
    MAKE_ARRAY      byte = iota
    MAKE_SLICE      byte = iota
    MAKE_LINKED     byte = iota
    MAKE_SET        byte = iota
    MAKE_MAP        byte = iota
    MAKE_STACK      byte = iota
    MAKE_QUEUE      byte = iota
    MAKE_DEQUE      byte = iota
    MAKE_SCOOP      byte = iota
    STORE_VAR       byte = iota
    LOAD_VAR        byte = iota
    UNARY_NOT       byte = iota
    BINARY_AND      byte = iota
    BINARY_OR       byte = iota
    BINARY_XOR      byte = iota
    BINARY_ADD      byte = iota
    BINARY_SUB      byte = iota
    BINARY_MUL      byte = iota
    BINARY_DIV      byte = iota
    BINARY_POW      byte = iota
    PRINT_OBJ       byte = iota
    PRINT_NEWLINE   byte = iota
    ABSOLUTE_JUMP   byte = iota
    RELATIVE_JUMP   byte = iota
    JUMP_IF_TRUE    byte = iota
    JUMP_IF_FALSE   byte = iota
    CALL_SCOOP      byte = iota
    RETURN          byte = iota
    POP             byte = iota
)
