/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "exec.hpp" header file for the Scoops Virtual Machine (SVM). 
 * It is used by "svm.cpp" to execute bytecode instructions.
 *
 * "exec" function ...
 *     1. Accepts a few arguments as pointers:
 *            - "running": a binary switch that gets switched off by the THE_END
 *              opcode. Yeah, it's not really that popular;
 *            - "error": another binary switch that gets switched on whenever
 *              SVM encounters an error during execution;
 *            - "ip": instruction pointer that represents the subscript of
 *              current instruction in ...
 *            - "memory": vector that holds all instructions;
 *            - "data": stack used by SVM to manipulate instances of ScpObj.
 *     2. Using the switch statement on current instruction's opcode, it 
 *        executes a function to which the opcode refers (all opcode functions
 *        are located in the "ops.hpp" header file, opcode definitions are
 *        hidden in "../Include/opcodes.hpp").
 *     3. In case it finds an unknown opcode, it calls the function called
 *        "bad_opcode" located in the "ops.hpp" with the rest of the opcode
 *         functions. That function simply switches the "error" switch on.
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

        case BINARY_OP:
            binary_op(error, ip, memory, data);
            break;
        
/*
        case COMPARE_OP:
            compare_op(error, ip, memory, data);
            break;
*/

        case PRINT_NEWLINE:
            print_newline(ip);
            break;
        
        default:
            bad_opcode(error);
    }
}

