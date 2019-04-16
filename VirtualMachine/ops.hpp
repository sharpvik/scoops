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
             std::stack<ScpObj*> &data) 
{
    memory.clear();
    while ( !data.empty() ) data.pop();
    running = false;
    printf("\nTHE_END\n");
}


// 1
void nop(unsigned int &ip) 
{ 
    ip++; 
    printf("NOP\n");
}


// 2
void push_const(bool &error, unsigned int &ip, std::vector<INSTRUCTION> &memory,
                std::stack<ScpObj*> &data)
{
    INSTRUCTION instruction = memory[ip];
    
    switch (instruction[1])
    {
        case 'b':
            data.push( new ScpBln(instruction[2]) );
            break;

        case 'c':
            data.push( new ScpChr(instruction[2]) );
            break;

        case 'n':
            {
                BYTE cast[8];
                for (int i = 2; i < 2 + 8; i++)
                    cast[i - 2] = instruction[i];
                double num = *reinterpret_cast<double*>(cast);

                data.push( new ScpNum(num) );
            }
            break;

        default:
            error = true;
    }

    ip++;
}


// 3
void emit_const(bool &error, unsigned int &ip,
                std::stack<ScpObj*> &data)
{
    if ( data.empty() )
    {
        printf("ERROR: Empty Stack. Cannot EMIT_CONST.\n");
        error = true;
        return;
    }

    ScpObj* ptr = data.top();
    std::string type = ptr->get_type();

    if (type == "bln")
        std::cout << ( ( (ScpBln*)ptr )->get_value() ? "true" : "false" );

    else if (type == "chr")
        std::cout << ( ( (ScpChr*)ptr )->get_value() );

    else if (type == "num")
        std::cout << ( ( (ScpNum*)ptr )->get_value() );

    else if (type == "str")
        std::cout << ( ( (ScpStr*)ptr )->get_value() );

    else 
    {
        printf("ERROR: Unknown Const Type. Cannot EMIT_CONST.\n");
        error = true;
        return;
    }
    
    delete ptr;
    data.pop();
    ip++;
}


// 6
void print_newline(unsigned int &ip)
{
    printf("\n");
    ip++;
}

