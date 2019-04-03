/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the main "sasm.cpp" file for the Scoops Assembler (SASM). 
 * As Scoops Virtual Machine (SVM) operates on bytecode it is hard to write
 * programs for it -- programming in bytes is not much fun -- it's almost like
 * coding in binary. Therefore, SASM has been created to provide an intermediate
 * stage between Scoops source code (.scp) files and Scoops bytecode (.scpb) 
 * files. SASM's file format is ".scpa" which stands for Scoops Assembly.
 *
 * This file relies on the "../svm/util.hpp" for opcode and function definitions
 * that -- as you can imagine -- this assemly language is based on.
 *
 *
 * Usage details:
 *     (Windows) CMD:
 *     sasm.exe <path to source file> <path to output file>
 *    
 *     (Mac OS / Linux) Terminal:
 *     ./sasm <path to source file> <path to output file>
 *
 *
 * Command structure:
 *     <opcode> (<operand>, <operand>, ...)*
 *     *part in the brackets is optional
 *
 *
 * Example program:
 *     # just a comment
 *     # this program prints the meaning of life
 *     PUSH_CONST 41
 *     PUSH_CONST 1
 *     BINARY_ADD
 *     EMIT_CONST
 *     PRINT_NEWLINE
 *     HALT
 *
 * 
 * Return codes and their meanings:
 *     0 = "OK"
 *     1 = "no filename specified"
 * 
 */
 


#include <stdio.h>
#include <string.h>
#include <sstream>
#include <vector>

#include "../svm/util.hpp"
#include "lex.hpp"
#include "err.hpp"
#include "gen.hpp"



void usage()
{
    std::string u = 
        "Usage details:\n "
        "\t(Windows) CMD:\n "
        "\t~$ sasm.exe <infile path>/<filename>.sasm <outfile path>/<filename>.svmb\n "
        "\n "
        "\t(Mac OS / Linux) Terminal:\n "
        "\t~$ vm <infile path>/<filename>.sasm <outfile path>/<filename>.svmb";
    std::cout << u << "\n";
}



int main(int argc, char* argv[])
{
    // READ FILE LINE BY LINE
    char* filename;
    // check if filename was specified
    if (argc < 2)
    {
        usage();
        return 1; // return code 1 = "no filename specified"
    }
    filename = argv[1];
    // check if file exists
    if ( !FileExists(filename) )
    {
        printf("File does not exist.\n");
        return 404; // return code 404 = "file not found"
    }
    // read line by line
    std::ifstream file(filename);
    std::string line;
    while ( std::getline(file, line) )
    {
        // RUN LEXER
        //
        
        
        // RUN ERROR CHECKER
        //
        
        
        // RUN BYTECODE GEN
        //
    }
    
    
    /* WRITE TO FILE
    // check if output filename was specified
    if (argc < 3)
    {
        // dump file into the current directory
    }
    else
    {
        // write to path specified
    } */
    
    
    return 0;
}