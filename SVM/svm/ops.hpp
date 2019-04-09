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

#include <stdio.h>
#include <stack>

#include "util.hpp"



// 1
char* push_const(char* ip, std::stack<OBJECT>* data)
{
    OBJECT o;
    o.type = *(ip + 1);
    switch (o.type) {
        case 'b':
            o.b = *(ip + 2) > 0;
            data->push(o);
            return ip + 3;
            break;
        
        case 'c':
            o.c = *(ip + 2);
            data->push(o); 
            return ip + 3;
            break;
            
        case 'n':
            o.n = *reinterpret_cast<double*>(ip + 2);
            data->push(o);
            return ip + 10;
            break;
    }
}


// 2
char* emit_const(char* ip, std::stack<OBJECT>* data) 
{ 
    OBJECT o = data->top();
    data->pop();
    
    switch (o.type) {
        case 'b':
            printf(o.b ? "true" : "false");
            break;
            
        case 'c':
            printf("%c", o.c);
            break;
            
        case 'n':
            printf("%f", o.n);
            break;
    }
    
    return ip + 1; 
}


// 3
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


// 4
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


// 5
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


// 6
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


// 7
char* binary_not(char* ip, std::stack<OBJECT>* data)
{
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = !a.b;
    data->push(o);
    return ip + 1;
}


// 8
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


// 9
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


// 10
char* binary_xor(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = a.b ^ b.b;
    data->push(o);
    return ip + 1;
}


// 11
char* print_newline(char* ip)
{
    printf("\n");
    return ip + 1;
}


// 12
char* equal_to(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = a.n == b.n;
    data->push(o);
    return ip + 1;
}


// 13
char* not_eq_to(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = a.n != b.n;
    data->push(o);
    return ip + 1;
}


// 14
char* grt_than(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = a.n > b.n;
    data->push(o);
    return ip + 1;
}


// 15
char* grt_eq_to(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = a.n >= b.n;
    data->push(o);
    return ip + 1;
}


// 16
char* less_than(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = a.n < b.n;
    data->push(o);
    return ip + 1;
}


// 17
char* less_eq_to(char* ip, std::stack<OBJECT>* data)
{
    OBJECT b = data->top(); data->pop();
    OBJECT a = data->top(); data->pop();
    OBJECT o;
    o.type = 'b';
    o.b = a.n <= b.n;
    data->push(o);
    return ip + 1;
}


// 18
char* nop(char* ip) { return ip + 1; }
