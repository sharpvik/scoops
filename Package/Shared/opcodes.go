package shared

/*
 * If you were to ever edit the opcodes down here, make sure that you
 *     1. Run scoops/Tools/assembly_instructions_update.py
 *     2. Update...
 *         │ scoops/Package/Assembly/assembler_test.go
 *         │ scoops/InputFiles/scpa/*
 */

const (                             // │ №   │ hex │ № operands │
    THE_END         uint8 = iota    // │ 0   │ x00 │ 0          │
    NOP             uint8 = iota    // │ 1   │ x01 │ 0          │
    LOAD_CONST      uint8 = iota    // │ 2   │ x02 │ 1          │
    MAKE_PRIMITIVE  uint8 = iota    // │ 3   │ x03 │ 1          │
    STORE_VAR       uint8 = iota    // │ 4   │ x04 │ 1          │
    LOAD_VAR        uint8 = iota    // │ 5   │ x05 │ 1          │
    BINARY_OP       uint8 = iota    // │ 6   │ x06 │ 1          │
    MAKE_COLLECTION uint8 = iota    // │ 7   │ x07 │ 2          │
    PRINT_OBJ       uint8 = iota    // │ 8   │ x08 │ 0          │
    PRINT_NEWLINE   uint8 = iota    // │ 9   │ x09 │ 0          │
    ABSOLUTE_JUMP   uint8 = iota    // │ 10  │ x0A │ 1          │
    RELATIVE_JUMP   uint8 = iota    // │ 11  │ x0B │ 1          │
    JUMP_IF_TRUE    uint8 = iota    // │ 12  │ x0C │ 1          │
    JUMP_IF_FALSE   uint8 = iota    // │ 13  │ x0D │ 1          │
    MAKE_SCOOP      uint8 = iota    // │ 14  │ x0E │ 2          │
    CALL_SCOOP      uint8 = iota    // │ 15  │ x0F │ 1          │
    RETURN          uint8 = iota    // │ 16  │ x10 │ 1          │
)
