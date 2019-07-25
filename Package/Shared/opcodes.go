package shared

/*
 * If you were to ever edit the opcodes down here, make sure that you
 *     1. Run scoops/Tools/assembly_instructions_update.py
 *     2. Possibly run scoops/Tools/enumerate_opcodes.py
 *     3. Update if necessary...
 *         │ scoops/Package/Assembly/assembler_test.go
 *         │ scoops/InputFiles/scpa/*
 */

const (
    END             byte = iota
    PUSH_CONST
    MAKE_BLN
    MAKE_BYTE
    MAKE_FLT
    MAKE_INT
    MAKE_NIL
    MAKE_RUNE
    MAKE_STRING
    STRING_CONCATENATE
    MAKE_SLICE
    SLICE_APPEND
    SLICE_GET_ITEM_BY_INDEX
    SLICE_POP
    MAKE_LINKED
    MAKE_SET
    MAKE_MAP
    MAKE_STACK
    MAKE_QUEUE
    MAKE_DEQUE
    MAKE_SCOOP
    PRINT_OBJ
    GET_TYPE
    GET_SIZE
    STORE_VAR
    LOAD_VAR
    UNARY_NOT
    BINARY_AND
    BINARY_OR
    BINARY_XOR
    BINARY_ADD
    BINARY_SUB
    BINARY_MUL
    BINARY_DIV
    BINARY_POW
    PRINT_NEWLINE
    ABSOLUTE_JUMP
    RELATIVE_JUMP
    JUMP_IF_TRUE
    JUMP_IF_FALSE
    CALL_SCOOP
    RETURN
    POP
)
