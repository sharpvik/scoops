package shared

/*
 * If you were to ever edit the opcodes down here, make sure that you
 *     1. Run scoops/Tools/assembly_instructions_update.py
 *     2. Update...
 *         │ scoops/Package/Assembly/assembler_test.go
 *         │ scoops/InputFiles/scpa/*
 */
const ( /* opcodes */
    THE_END uint8 = iota            // 0
    NOP uint8 = iota                // 1
    LOAD_BYTES uint8 = iota         // 2
    PRINT_OBJECT uint8 = iota       // 3
    ARITHMETIC_OP uint8 = iota      // 4
    BINARY_OP uint8 = iota          // 5
    COMPARE_OP uint8 = iota         // 6
    PRINT_NEWLINE uint8 = iota      // 7
)
