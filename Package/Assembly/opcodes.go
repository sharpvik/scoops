package assembly

import "github.com/sharpvik/scoops/Package/Shared"

// Based on scoops/Package/Shared/opcodes.go/const(opcodes)
var OpcodeMap = map[string]uint8{
    "END": shared.END,
    "PUSH_CONST": shared.PUSH_CONST,
    "MAKE_BLN": shared.MAKE_BLN,
    "MAKE_FLT": shared.MAKE_FLT,
    "MAKE_INT": shared.MAKE_INT,
    "MAKE_NIL": shared.MAKE_NIL,
    "MAKE_RUNE": shared.MAKE_RUNE,
    "MAKE_STRING": shared.MAKE_STRING,
    "STRING_CONCATENATE": shared.STRING_CONCATENATE,
    "MAKE_SLICE": shared.MAKE_SLICE,
    "SLICE_APPEND": shared.SLICE_APPEND,
    "SLICE_GET_ITEM_BY_INDEX": shared.SLICE_GET_ITEM_BY_INDEX,
    "SLICE_POP": shared.SLICE_POP,
    "MAKE_LINKED": shared.MAKE_LINKED,
    "MAKE_SET": shared.MAKE_SET,
    "MAKE_MAP": shared.MAKE_MAP,
    "MAKE_STACK": shared.MAKE_STACK,
    "MAKE_QUEUE": shared.MAKE_QUEUE,
    "MAKE_DEQUE": shared.MAKE_DEQUE,
    "MAKE_SCOOP": shared.MAKE_SCOOP,
    "CLONE_OBJ": shared.CLONE_OBJ,
    "PRINT_OBJ": shared.PRINT_OBJ,
    "GET_TYPE": shared.GET_TYPE,
    "GET_SIZE": shared.GET_SIZE,
    "STORE_VAR": shared.STORE_VAR,
    "LOAD_VAR": shared.LOAD_VAR,
    "UNARY_NOT": shared.UNARY_NOT,
    "BINARY_AND": shared.BINARY_AND,
    "BINARY_OR": shared.BINARY_OR,
    "BINARY_XOR": shared.BINARY_XOR,
    "BINARY_ADD": shared.BINARY_ADD,
    "BINARY_SUB": shared.BINARY_SUB,
    "BINARY_MUL": shared.BINARY_MUL,
    "BINARY_DIV": shared.BINARY_DIV,
    "BINARY_POW": shared.BINARY_POW,
    "PRINT_NEWLINE": shared.PRINT_NEWLINE,
    "ABSOLUTE_JUMP": shared.ABSOLUTE_JUMP,
    "RELATIVE_JUMP": shared.RELATIVE_JUMP,
    "JUMP_IF_TRUE": shared.JUMP_IF_TRUE,
    "JUMP_IF_FALSE": shared.JUMP_IF_FALSE,
    "CALL_SCOOP": shared.CALL_SCOOP,
    "RETURN": shared.RETURN,
    "POP": shared.POP,
}
