/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "ops.hpp" header file for the Scoops Virtual Machine (SVM). 
 * It is used by "exec.hpp" as a library of opcode related functioons. Every
 * function below has a unique opcode (integer that is commented out one line
 * above its declaration) that trigers its execution.
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
    unsigned char type = instruction[1];
    
    switch (type)
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
void emit_const(bool &error, unsigned int &ip, std::stack<ScpObj*> &data)
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


// 4
void binary_op(bool &error, unsigned int &ip, std::vector<INSTRUCTION> &memory,
               std::stack<ScpObj*> &data)
{
    INSTRUCTION instruction = memory[ip];
    
    unsigned char sign = instruction[1];
    
    if (sign == '!')
    {
        std::string type = data.top()->get_type();
        if (type != "bln")
        {
            printf("ERROR: Incorrect data type for BINARY_OP.\n");
            error = true;
            return;
        }

        ScpBln* ptr = (ScpBln*)data.top();
        ptr->binary_not();
    }

    else if (sign == '&' || sign == '|' || sign == '^')
    {
        ScpObj* second = data.top(); data.pop();
        ScpObj* first  = data.top(); data.pop();
        std::string ftype = first->get_type();
        std::string stype = second->get_type();
        if (ftype != "bln" || stype != "bln")
        {
            printf("ERROR: Incorrect data type for BINARY_OP.\n");
            error = true;
            return;
        }
        ScpBln* fptr = (ScpBln*)first;
        ScpBln* sptr = (ScpBln*)second;
        bool new_value;

        switch (sign)
        {
            case '&':
                new_value = fptr->get_value() && sptr->get_value();
                break;

            case '|':
                new_value = fptr->get_value() || sptr->get_value();
                break;

            case '^':
                new_value = fptr->get_value() ^ sptr->get_value();
                break;
        }

        delete fptr; delete sptr;
        ScpBln* new_ptr = new ScpBln(new_value);
        data.push(new_ptr);
    }

    else if (sign == '+' || sign == '-' || sign == '*' || sign == '/')
    {
        ScpObj* second = data.top(); data.pop();
        ScpObj* first  = data.top(); data.pop();
        std::string ftype = first->get_type();
        std::string stype = second->get_type();
        if (ftype != "num" || stype != "num")
        {
            printf("ERROR: Incorrect data type for BINARY_OP.\n");
            error = true;
            return;
        }
        ScpNum* fptr = (ScpNum*)first;
        ScpNum* sptr = (ScpNum*)second;
        double new_value;

        switch (sign)
        {
            case '+':
                new_value = fptr->get_value() + sptr->get_value();
                break;

            case '-':
                new_value = fptr->get_value() - sptr->get_value();
                break;

            case '*':
                new_value = fptr->get_value() * sptr->get_value();
                break;

            case '/':
                new_value = fptr->get_value() / sptr->get_value();
                break;
        }

        delete fptr; delete sptr;
        ScpNum* new_ptr = new ScpNum(new_value);
        data.push(new_ptr);
    }

    else
    {
        error = true;
        printf("ERROR: Invalid sign for BINARY_OP.\n");
        return;
    }

    ip++;
}


// 5
//


// 6
void print_newline(unsigned int &ip)
{
    printf("\n");
    ip++;
}

