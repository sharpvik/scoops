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
    PUSH_CONST      byte = iota	// 1
    MAKE_BLN        byte = iota	// 2
    MAKE_BYTE       byte = iota	// 3
    MAKE_FLT        byte = iota	// 4
    MAKE_INT        byte = iota	// 5
    MAKE_NIL        byte = iota	// 6
    MAKE_RUNE       byte = iota	// 7
    MAKE_STR        byte = iota	// 8
    MAKE_ARRAY      byte = iota	// 9
    MAKE_SLICE      byte = iota	// 10
    MAKE_LINKED     byte = iota	// 11
    MAKE_SET        byte = iota	// 12
    MAKE_MAP        byte = iota	// 13
    MAKE_STACK      byte = iota	// 14
    MAKE_QUEUE      byte = iota	// 15
    MAKE_DEQUE      byte = iota	// 16
    MAKE_SCOOP      byte = iota	// 17
    STORE_VAR       byte = iota	// 18
    LOAD_VAR        byte = iota	// 19
    UNARY_NOT       byte = iota	// 20
    BINARY_AND      byte = iota	// 21
    BINARY_OR       byte = iota	// 22
    BINARY_XOR      byte = iota	// 23
    BINARY_ADD      byte = iota	// 24
    BINARY_SUB      byte = iota	// 25
    BINARY_MUL      byte = iota	// 26
    BINARY_DIV      byte = iota	// 27
    BINARY_POW      byte = iota	// 28
    PRINT_OBJ       byte = iota	// 29
    PRINT_NEWLINE   byte = iota	// 30
    ABSOLUTE_JUMP   byte = iota	// 31
    RELATIVE_JUMP   byte = iota	// 32
    JUMP_IF_TRUE    byte = iota	// 33
    JUMP_IF_FALSE   byte = iota	// 34
    CALL_SCOOP      byte = iota	// 35
    RETURN          byte = iota	// 36
    POP             byte = iota	// 37
)
