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
 * that -- as you can imagine -- this assemly language is based on. It also
 * relies on the header (.hpp) files in this directory.
 *
 * What SASM does:
 *     1. Reads input ".scpa" file line by line;
 *     2. Runs the "lex" function (from "lex.hpp") -- passing each line as an
 *        argument -- that returns a vector of TOKENs (TOKEN type 
 *        is defined in "lex.hpp");
 *     3. Error checking based on lexer's output -- namely checking for TOKENs
 *        with "error" type;
 *     4. Bytecode generation based on the vector of TOKENs.
 *     5. Writing bytecode to the output file specified or to file in the local
 *        directory. Output file is of ".scpb" format, which stands for Scoops
 *        bytecode. Bytecode files will be used by SVM.
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
 *     PUSH_CONST #41
 *     PUSH_CONST #1
 *     BINARY_ADD
 *     EMIT_CONST
 *     PRINT_NEWLINE
 *     HALT
 *
 * 
 * Return codes and their meanings:
 *     0   = "OK"
 *     1   = "no filename specified"
 *     2   = "syntax error"
 *     404 = "file not found"
 * 
 */
 


#include <stdio.h>
#include <string.h>
#include <sstream>
#include <vector>

#include "../svm/util.hpp"
#include "lex.hpp"
#include "gen.hpp"



// UTILITY FUNCTIONS
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


bool error_detected(std::vector<TOKEN> toks)
{
    int toks_size = toks.size();
    for (int i = 0; i < toks_size; i++)
    {
        TOKEN t = toks[i];
        if (t.type == "error") return true;
    }
    return false;
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
    unsigned int line_count = 1;
    std::ifstream file(filename);
    std::string line;
    while ( std::getline(file, line) )
    {
        // SKIP EMPTY LINES AND LINES WITH COMMENTS
        if ( line[0] == 0 || line.substr(0, 2) == "//" )
        {
            //printf("Line %d ignored.\n", line_count); // debug
            line_count++;
            continue;
        }
        
        
        // RUN LEXER
        std::vector<TOKEN> toks = lex(line + " ");
        /*----------------------------- LEX DEBUG ----------------------------*/
        std::cout << line << "\n"; 
        for (int i = 0; i < toks.size(); i++)
        {
            TOKEN t = toks[i];
            std::cout << t.type << " "; 
        } std::cout << "\n";
        /*--------------------------------------------------------------------*/
        
        
        // RUN ERROR CHECKER
        if ( error_detected(toks) )
        {
            printf("\nSYNTAX ERROR on line %d.\n", line_count);
            printf("Assembler exited with non-zero return value.\n");
            return 2;
        }
        
        
        // RUN BYTECODE GEN
        //
        
        
        line_count++;
    }
    
    
    /*----------------------------- WRITE TO FILE -----------------------------/
    // check if output filename was specified
    if (argc < 3)
    {
        // dump file into the current directory
    }
    else
    {
        // write to path specified
    } 
    /*------------------------------------------------------------------------*/
    
    
    return 0; // return code 0 = "OK"
}
