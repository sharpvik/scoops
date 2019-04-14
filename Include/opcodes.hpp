/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "opcodes.hpp" file. Guess what it contains!
 *
 */



#pragma once



enum Opcode {
/*-------------------- byte | # of operand(s) | # byte(s) for operands -------*/
    THE_END,        // 0x00 | 0               | 0
    NOP,            // 0x01 | 0               | 0
    PUSH_CONST,     // 0x02 | 1               | 2 - 9
    EMIT_CONST,     // 0x03 | 0               | 0
    BINARY_OP,      // 0x04 | 1               | 1
    COMPARE_OP,     // 0x05 | 2               | 0
    PRINT_NEWLINE,  // 0x06 | 0               | 0
    INSTRUCTION_END = 10,
};
