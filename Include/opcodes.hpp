/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "opcodes.hpp" file. Guess what it contains!
 *
 */



#pragma once



enum Opcode {
/*-------------------- ord # | # of operand(s) | # byte(s) for operands ------*/
    THE_END,        // 0     | 0               | 0
    NOP,            // 1     | 0               | 0
    PUSH_CONST,     // 2     | 1               | 2 - 9
    EMIT_CONST,     // 3     | 0               | 0
    BINARY_OP,      // 4     | 1               | 1
    COMPARE_OP,     // 5     | 2               | 0
    PRINT_NEWLINE,  // 6     | 0               | 0
    INSTRUCTION_END = 10,
};
