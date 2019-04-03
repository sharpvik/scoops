/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "lex.hpp" header file for the Scoops Assembler (SASM). 
 * It is used by "sasm.cpp" as a header file for function definitions that are 
 * required for lexical analysis of SASM instructions. As you can see, it also
 * has a few vectors that hold values required to check lexems' type.
 *
 * To operate correctly, vector "ops" must always contain the latest most 
 * relevant list of Scoops Virtual Machine (SVM) opcodes and in the right order!
 * If any opcodes were to be added, changed or altered, this change must be
 * reflected on this file -- and more specifically on the "ops" vector.
 *
 */
 
 

#include <iostream>
#include <vector>
#include <string>



// LEXEMES
std::vector<std::string> ops
{
    "HALT", "PUSH_CONST", "EMIT_CONST", "BINARY_ADD", "BINARY_SUB", 
    "BINARY_MUL", "BINARY_DIV", "BINARY_NOT", "BINARY_AND", "BINARY_OR", 
    "BINARY_XOR", "PRINT_NEWLINE", "EQUAL_TO", "NOT_EQ_TO", "GRT_THAN", 
    "GRT_EQ_TO", "LESS_THAN", "LESS_EQ_TO",
};

std::vector<char> digits{ '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' };

std::vector<std::string> bools{ "true", "false" };



// UTILITY FUNCTIONS
bool lex_find(std::vector<std::string> arr, std::string needle)
{
    int arr_size = arr.size();
    for (int i = 0; i < arr_size; i++)
        if (arr[i] == needle) return i;
    return -1;
}
