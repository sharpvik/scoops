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

#include <stack>

#include "util.hpp"
#include "ops.hpp"



char* exec(char *ip, std::stack<OBJECT> *data)
{
    switch (*ip)
    {
        case PUSH_BLN:
            return push_bln(ip, data);
            break;
        
        case PUSH_CHR:
            return push_chr(ip, data);
            break;
            
        case PUSH_NUM:
            return push_num(ip, data);
            break;
            
        case EMIT_CONST:
            return emit_const(ip, data);
            break;
        
        default:
            return nop(ip);
    }
}