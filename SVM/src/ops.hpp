/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "ops.hpp" header file for the Scoops Virtual Machine (SVM). 
 * It is used by "exec.hpp" as a library of opcode related functioons. Every
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

#include "util.hpp"



// anything excecpt codes below
char* nop(char* ip) { return ip + 1; }


// 1
char* emit_const(char* ip, std::stack<OBJECT>* data) 
{ 
    OBJECT o = data->top();
    data->pop();
    
    switch (o.type) {
        case 'b': // bln
            printf(o.b ? "true" : "false");
            break;
            
        case 'c': // chr
            printf("%c", o.c);
            break;
            
        case 'n': // num
            printf("%f", o.n);
            break;
    }
    
    return ip + 1; 
}


// 2
// receives next byte as operand (converts it to 'bool')
char* push_bln(char* ip, std::stack<OBJECT>* data)
{
    OBJECT o;
    o.type = 'b';
    o.b = *(ip + 1) > 0;
    data->push(o);
    return ip + 2;
}


// 3
// receives next byte as operand (converts it to 'char')
char* push_chr(char* ip, std::stack<OBJECT>* data) 
{
    OBJECT o;
    o.type = 'c';
    o.c = *(ip + 1);
    data->push(o); 
    return ip + 2;
}


// 4
// receives next 8 bytes as operand (converts them to 'double')
char* push_num(char* ip, std::stack<OBJECT>* data)
{
    OBJECT o;
    o.type = 'n';
    o.n = *reinterpret_cast<double*>(ip + 1);
    data->push(o);
    return ip + 9;
}


// 5
char* binary_add(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'n';
    o.n = a.n + b.n;
    data->push(o);
    return ip + 1;
}


// 6
char* binary_sub(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'n';
    o.n = a.n - b.n;
    data->push(o);
    return ip + 1;
}


// 7
char* binary_mul(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'n';
    o.n = a.n * b.n;
    data->push(o);
    return ip + 1;
}


// 8
char* binary_div(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'n';
    o.n = a.n / b.n;
    data->push(o);
    return ip + 1;
}


// 9
char* binary_not(char* ip, std::stack<OBJECT>* data)
{
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = !a.b;
    data->push(o);
    return ip + 1;
}


// 10
char* binary_and(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = a.b && b.b;
    data->push(o);
    return ip + 1;
}


// 11
char* binary_or(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = a.b || b.b;
    data->push(o);
    return ip + 1;
}


// 12
char* binary_xor(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = a.b && !b.b || !a.b && b.b;
    data->push(o);
    return ip + 1;
}


// 13
char* print_newline(char* ip)
{
    printf("\n");
    return ip + 1;
}
