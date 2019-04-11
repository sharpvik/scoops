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
        
        case BINARY_OP:
            return binary_op(ip, data);
            break; 
        
        case PRINT_NEWLINE:
            return print_newline(ip);
            break;
        
        case EQUAL_TO:
            return equal_to(ip, data);
            break;
        
        case NOT_EQ_TO:
            return not_eq_to(ip, data);
            break;
        
        case GRT_THAN:
            return grt_than(ip, data);
            break;
        
        case GRT_EQ_TO:
            return grt_eq_to(ip, data);
            break;
        
        case LESS_THAN:
            return less_than(ip, data);
            break;
        
        case LESS_EQ_TO:
            return less_eq_to(ip, data);
            break;
        
        case NOP:
            return nop(ip);
            break;
        
        default:
            return nop(ip);
    }
}
