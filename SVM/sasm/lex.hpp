/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "lex.hpp" header file for the Scoops Assembler (SASM). 
 * It is used by "sasm.cpp" as a header file for function definitions that are 
 * required for lexical analysis of SASM instructions. As you can see, it also
 * has a few vectors that hold values required to check lexems' type.
 *
 * To operate correctly, vector "ops" must always contain the latest most 
 * relevant list of Scoops Virtual Machine (SVM) opcodes and in the right order!
 * If any opcodes were to be added, changed or altered, _this change must be
 * reflected on _this file -- and more specifically on the "ops" vector.
 *
 * The "lex" function is a state machine that uses the "buffer" variable 
 * to accumulate characters it reads from the "line" string and produce a vector
 * of TOKENs (TOKEN type is defined below).
 *
 */



#pragma once

#include <iostream>
#include <vector>
#include <string>
#include <sstream>



// TYPE DEFINITIONS
typedef struct TOKEN_t 
{
    std::string type;
    
    union {
        bool    bln_val;
        char    chr_val;
        double  dbl_val;
        char    int_val;
        //char*   str_val;
    };
} TOKEN;



// LEXEMES
std::vector<std::string> ops
{
    "HALT", "PUSH_CONST", "EMIT_CONST", "BINARY_ADD", "BINARY_SUB", 
    "BINARY_MUL", "BINARY_DIV", "BINARY_NOT", "BINARY_AND", "BINARY_OR", 
    "BINARY_XOR", "PRINT_NEWLINE", "EQUAL_TO", "NOT_EQ_TO", "GRT_THAN", 
    "GRT_EQ_TO", "LESS_THAN", "LESS_EQ_TO",
};



// STATES
enum State {
    NONE,           // 0
    OPCODE_READ,    // 1
    CHAR_READ,      // 2
    NUM_READ,       // 3
    BOOL_READ,      // 4
};



// UTILITY FUNCTIONS
char lex_find(std::vector<std::string> arr, std::string needle)
{
    int arr_size = arr.size();
    for (int i = 0; i < arr_size; i++)
        if (arr[i] == needle) return i;
    return -1;
}

bool num_check(std::string buffer)
{
    int buffer_length = buffer.size();
    for (int i = 0; i < buffer_length; i++)
    {
        char _this = buffer[i];
        if (_this < '0' || _this > '9' && _this != '.') return false;
    }
    return true;
}



// LEXER FUNCTION
std::vector<TOKEN> lex(std::string line)
{
    State current_state = OPCODE_READ;
    std::string buffer;
    std::vector<TOKEN> toks;
    
    int line_length = line.length();
    for (int i = 0; i < line_length; i++)
    {
        char _this = line[i];
        //std::cout << current_state << " " << _this << "\n"; // debug
        switch (current_state) 
        {
            case NONE: // decide on the next state
                switch (_this) 
                {
                    case 'b':
                        current_state = BOOL_READ;
                        break;
                        
                    case '\'':
                        current_state = CHAR_READ;
                        break;
                        
                    case '#':
                        current_state = NUM_READ;
                        break;
                        
                    default:
                        current_state = NONE;
                }
                break;
                
            case OPCODE_READ:
                if (_this == ' ') // done reading opcode
                {
                    TOKEN t;
                    char ord = lex_find(ops, buffer);
                    t.type = ord != -1 ? "opcode" : "error";
                    t.int_val = ord != -1 ? ord : -1 ;
                    toks.push_back(t);
                    buffer = "";
                    current_state = NONE;
                }
                else buffer += _this; // continue reading opcode
                break;
                
            case BOOL_READ:
                if (_this == ' ') // done reading bool
                {
                    TOKEN t;
                    int buffer_length = buffer.length();
                    t.type = buffer_length == 1 ? "bool" : "error";
                    t.bln_val = buffer_length != 1 ? 1 : buffer[0] == '1';
                    toks.push_back(t);
                    buffer = "";
                    current_state = NONE;
                }
                else buffer += _this; // continue reading bool
                break;
                
            case CHAR_READ:
                if (_this == '\'') // done reading char
                {
                    TOKEN t;
                    int buffer_length = buffer.length();
                    t.type = buffer_length == 1 ? "char" : "error";
                    t.chr_val = buffer_length == 1 ? buffer[0] : 'e';
                    toks.push_back(t);
                    buffer = "";
                    current_state = NONE;
                }
                else buffer += _this; // continue reading char
                break;
                
            case NUM_READ:
                if (_this == ' ') // done reading num
                {
                    TOKEN t;
                    bool num_correct = num_check(buffer);
                    t.type = num_correct ? "num" : "error";
                    t.dbl_val = num_correct ? std::stod(buffer) : 1;
                    toks.push_back(t);
                    buffer = "";
                    current_state = NONE;
                }
                else buffer += _this; // continue reading num
                break;
        }
    }
    
    return toks;
}
