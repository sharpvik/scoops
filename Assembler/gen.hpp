/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "gen.hpp" header file for the Scoops Assembler (SASM). 
 * It is used by "sasm.cpp" as a header file for function definitions that are 
 * required for bytecode generation. 
 *
 */



#pragma once

#include <vector>
#include <stdio.h>


 
// UTILITY FUNCTIONS
unsigned char* double_to_bytes(double &n)
{ return reinterpret_cast<unsigned char*>(&n); }



bool bytecode_gen(std::vector<TOKEN> toks, std::vector<unsigned char>* buffer)
{
    int toks_size = toks.size();
    for (int i = 0; i < toks_size; i++)
    {
        TOKEN _this = toks[i];
        if (_this.type == "opcode")
            buffer->push_back( (unsigned char)_this.int_val );
        else if (_this.type == "bool")
        {
            buffer->push_back( (unsigned char)'b' );
            buffer->push_back( (unsigned char)_this.bln_val );
        }
        else if (_this.type == "char")
        {
            buffer->push_back( (unsigned char)'c' );
            buffer->push_back( (unsigned char)_this.chr_val );
        }    
        else if (_this.type == "num")
        {
            unsigned char* byted_num = double_to_bytes(_this.dbl_val);
            buffer->push_back( (unsigned char)'n' );
            for (int j = 0; j < 8; j++)
                buffer->push_back(byted_num[j]);
        }
        else return false;
    }
    return true;
}