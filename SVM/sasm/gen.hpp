/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "gen.hpp" header file for the Scoops Assembler (SASM). 
 * It is used by "sasm.cpp" as a header file for function definitions that are 
 * required for bytecode generation. 
 *
 */



#pragma once


 
// UTILITY FUNCTIONS
unsigned char* double_to_bytes(double &n)
{ return reinterpret_cast<unsigned char*>(&n); }