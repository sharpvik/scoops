/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "util.hpp" header file for the Scoops Virtual Machine (SVM). 
 * It is used as a header file for function and type definitions that are shared
 * between multiple files in this directory. Thus, you can see many other files
 * including it, while "util.hpp" itself mostly uses just the standard library.
 *
 * P.S. All SVM opcodes are defined here.
 *
 */
 


#pragma once

#include <iostream>
#include <string>
#include <fstream>



// OPCODES
enum Opcode {
    HALT,           // 0
    PUSH_CONST,     // 1
    EMIT_CONST,     // 2
    BINARY_ADD,     // 3
    BINARY_SUB,     // 4
    BINARY_MUL,     // 5
    BINARY_DIV,     // 6
    BINARY_NOT,     // 7
    BINARY_AND,     // 8
    BINARY_OR,      // 9
    BINARY_XOR,     // 10
    PRINT_NEWLINE,  // 11
    EQUAL_TO,       // 12
    NOT_EQ_TO,      // 13
    GRT_THAN,       // 14
    GRT_EQ_TO,      // 15
    LESS_THAN,      // 16
    LESS_EQ_TO,     // 17
    NOP,            // 18
};



// TYPE DECLARATIONS
typedef struct OBJECT_t 
{
    char type;
    
    union {
        bool    b;
        char    c;
        double  n;
        void*   p;
    };
} OBJECT;



// FILE RELATED FUNCTIONS
bool FileExists(char* filename)
{
    std::ifstream file(filename);
    return file.good();
}


char* ReadFile(char* filename)
{
    // open file
    std::ifstream file(filename);

    // get length of file
    file.seekg(0, file.end);
    size_t length = file.tellg();
    file.seekg(0, file.beg);

    // read file and return
    char* buffer;
    file.read(buffer, length);
    return buffer;
}
