/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "exec.hpp" header file for the Scoops Virtual Machine (SVM). 
 * It is used by "svm.cpp" to execute bytecode instructions.
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
#include "ops.hpp"



void exec(bool &running, bool &error, unsigned int &ip, 
          std::vector<INSTRUCTION> &memory, std::stack<ScpObj*> &data)
{
    INSTRUCTION instruction = memory[ip];
    BYTE opcode = instruction[0];
    //printf("%d ", opcode);
    switch (opcode)
    {
        case THE_END:
            the_end(running, memory, data);
            break;
        
        case NOP:
            nop(ip);
            break;

        case EMIT_CONST:
            emit_const(error, ip, data);
            break;
        
        case PUSH_CONST:
            push_const(error, ip, memory, data);
            break;

/*
        case BINARY_OP:
            binary_op(ip, memory, data);
            break;
        
        case COMPARE_OP:
            compare_op(ip, memory, data);
            break;
*/

        case PRINT_NEWLINE:
            print_newline(ip);
            break;
        
        default:
            bad_opcode(error);
    }
}

