package assembly

import "github.com/sharpvik/scoops/Package/Shared"

// Based on scoops/Package/Shared/opcodes.go/const(opcodes)
var OpcodeMap = map[string]uint8{
    "THE_END": shared.THE_END,
    "NOP": shared.NOP,
    "LOAD_CONST": shared.LOAD_CONST,
    "MAKE_PRIMITIVE": shared.MAKE_PRIMITIVE,
    "STORE_VAR": shared.STORE_VAR,
    "LOAD_VAR": shared.LOAD_VAR,
    "BINARY_OP": shared.BINARY_OP,
    "MAKE_COLLECTION": shared.MAKE_COLLECTION,
    "PRINT_OBJ": shared.PRINT_OBJ,
    "PRINT_NEWLINE": shared.PRINT_NEWLINE,
    "ABSOLUTE_JUMP": shared.ABSOLUTE_JUMP,
    "RELATIVE_JUMP": shared.RELATIVE_JUMP,
    "JUMP_IF_TRUE": shared.JUMP_IF_TRUE,
    "JUMP_IF_FALSE": shared.JUMP_IF_FALSE,
    "MAKE_SCOOP": shared.MAKE_SCOOP,
    "CALL_SCOOP": shared.CALL_SCOOP,
    "RETURN": shared.RETURN,
}
