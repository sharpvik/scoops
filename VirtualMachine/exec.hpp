/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "exec.hpp" header file for the Scoops Virtual Machine (SVM). 
 * It is used by "vm.cpp" to execute bytecode instructions.
 *
 * "exec" function ...
 *     1. Accepts instruction pointer (*ip) and pointer to the data stack 
 *        (*data) as its arguments;
 *     2. Using the switch statement with *ip, it executes a function to which
 *        the opcode refers (all opcode functions are located in
 *        the "ops.hpp" header file, preprocessor define statements for opcodes
 *        can be found in the "util.hpp" file);
 *     2. Returns new instruction pointer altered by the opcode 
 *        function that it executed (for details on how and why opcode functions
 *        alter the instruction pointer see "ops.hpp" file).
 *
 */
 


#pragma once



#include "util.hpp"
//#include "ops.hpp"



unsigned int exec(unsigned int ip, std::vector<INSTRUCTION>* memory, 
                  std::stack<OBJECT>* data)
{
    BYTE opcode = instruction[0];
    switch (opcode)
    {
        case EMIT_CONST:
            return emit_const(ip, data);
        
        case PUSH_CONST:
            return push_const(ip, memory, data);
        
        case BINARY_OP:
            return binary_op(ip, memory, data);
        
        case COMPARE_OP:
            return compare_op(ip, memory, data);
        
        case PRINT_NEWLINE:
            return print_newline(ip);

        case NOP:
            return nop(ip);
        
        default:
            return bad_opcode(ip, memory, data);
    }
}

