/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "ops.hpp" header file for the Scoops Virtual Machine (SVM). 
 * It is used by "exec.hpp" as a library of opcode related functioons. Every
 * function below has a unique opcode (integer that is commented out one line
 * above its declaration) that trigers its execution.
 *
 * For each opcode below, let R be the number of bytes it uses as operand(s).
 * Then, the pointer value (ip) that it returns must be increased by (R + 1).
 * This is done to keep track of the instruction pointer (*ip) shift due to
 * different amounts and sizes of operands required and used by each opcode.
 *
 */



#pragma once



#include "util.hpp"



void bad_opcode(bool &error) 
{ 
    error = true; 
    printf("BAD_OPCODE\n");
}


// 0
void the_end(bool &running, std::vector<INSTRUCTION> &memory,
             std::stack<ScpObj> &data) 
{
    memory.clear();
    while ( !data.empty() ) data.pop();
    running = false;
    printf("THE_END\n");
}


// 1
void nop(unsigned int &ip) 
{ 
    ip++; 
    printf("NOP\n");
}


// 6
void print_newline(unsigned int &ip)
{
    printf("\n");
    ip++;
}

