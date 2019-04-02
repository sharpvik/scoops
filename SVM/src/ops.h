/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "ops.h" header file for the Scoops Virtual Machine (SVM). 
 * It is used by "exec.h" as a library of opcode related functioons. Every
 * function below has a unique opcode (integer that is commented out one line
 * above its declaration) that trigers its execution.
 *
 * For each opcode below, let R be the number of operands it uses.
 * Then, the pointer value (ip) that it returns must be increased by (R + 1).
 * This is done to keep track of the instruction pointer (*ip) shift due to
 * different amounts of operands required and used by each opcode.
 *
 */



#pragma once

#include <stdio.h>
#include <stack>

#include "util.h"



// anything excecpt codes below
char * nop(char *ip) { return ip + 1; }


// 1
char * push_char(char *ip, std::stack<OBJECT> *data) 
{
    OBJECT o;
    o.type = "chr";
    o.c = *(ip + 1);
    data->push(o); 
    return ip + 2;
}


// 2
char * emit_char(char *ip, std::stack<OBJECT> *data) 
{ 
    OBJECT o = data->top();
    data->pop();
    printf("%c", o.c);
    return ip + 1; 
}