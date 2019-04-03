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
        case EMIT_CONST:
            return emit_const(ip, data);
            break;
            
        case PUSH_CONST:
            return push_const(ip, data);
            break;
        
        case BINARY_ADD:
            return binary_add(ip, data);
            break;
        
        case BINARY_SUB:
            return binary_sub(ip, data);
            break;
        
        case BINARY_MUL:
            return binary_mul(ip, data);
            break;
        
        case BINARY_DIV:
            return binary_div(ip, data);
            break;
        
        case BINARY_NOT:
            return binary_not(ip, data);
            break;
        
        case BINARY_AND:
            return binary_and(ip, data);
            break;
        
        case BINARY_OR:
            return binary_or(ip, data);
            break;
        
        case BINARY_XOR:
            return binary_xor(ip, data);
            break;
        
        case PRINT_NEWLINE:
            return print_newline(ip);
            break;
        
        default:
            return nop(ip);
    }
}
