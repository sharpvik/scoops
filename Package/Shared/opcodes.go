package shared

/*
 * If you were to ever edit the opcodes down here, make sure that you
 *     1. Run scoops/Tools/assembly_instructions_update.py
 *     2. Update...
 *         │ scoops/Package/Assembly/assembler_test.go
 *         │ scoops/InputFiles/scpa/*
 */

const (
    END             byte = iota	// 0
    NOP             byte = iota	// 1
    PUSH_CONST      byte = iota	// 2
    MAKE_BLN        byte = iota	// 3
    MAKE_BYTE       byte = iota	// 4
    MAKE_FLT        byte = iota	// 5
    MAKE_INT        byte = iota	// 6
    MAKE_NIL        byte = iota	// 7
    MAKE_RUNE       byte = iota	// 8
    MAKE_STR        byte = iota	// 9
    MAKE_ARRAY      byte = iota	// 10
    MAKE_SLICE      byte = iota	// 11
    MAKE_LINKED     byte = iota	// 12
    MAKE_SET        byte = iota	// 13
    MAKE_MAP        byte = iota	// 14
    MAKE_STACK      byte = iota	// 15
    MAKE_QUEUE      byte = iota	// 16
    MAKE_DEQUE      byte = iota	// 17
    MAKE_SCOOP      byte = iota	// 18
    STORE_VAR       byte = iota	// 19
    LOAD_VAR        byte = iota	// 20
    UNARY_NOT       byte = iota	// 21
    BINARY_AND      byte = iota	// 22
    BINARY_OR       byte = iota	// 23
    BINARY_XOR      byte = iota	// 24
    BINARY_ADD      byte = iota	// 25
    BINARY_SUB      byte = iota	// 26
    BINARY_MUL      byte = iota	// 27
    BINARY_DIV      byte = iota	// 28
    BINARY_POW      byte = iota	// 29
    PRINT_OBJ       byte = iota	// 30
    PRINT_NEWLINE   byte = iota	// 31
    ABSOLUTE_JUMP   byte = iota	// 32
    RELATIVE_JUMP   byte = iota	// 33
    JUMP_IF_TRUE    byte = iota	// 34
    JUMP_IF_FALSE   byte = iota	// 35
    CALL_SCOOP      byte = iota	// 36
    RETURN          byte = iota	// 37
    POP             byte = iota	// 38
)
