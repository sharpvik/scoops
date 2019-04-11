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
 *     3. Syntactic analysis and error checking based on lexer's output -- 
 *        namely checking for TOKENs with "error" type;
 *     4. Semantic analysis and error checking based on the number of operands
 *        required for each opcode;
 *     5. Bytecode generation based on the vector of TOKENs ("bytecode_gen"
 *        functions is defined in "gen.hpp");
 *     6. Writes bytecode to the output file specified or to file in the local
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
 *     3   = "semantic error"
 *     4   = "bytecode generation error"
 *     5   = "file open error"
 *     404 = "file not found"
 * 
 */
 


#include <stdio.h>
#include <string.h>
#include <sstream>
#include <vector>
#include <fstream>

#include "../svm/util.hpp"
#include "lex.hpp"
#include "gen.hpp"



// GLOBAL VARIABLES
/*
 * The "number_of_operands" array stores the number of operands needed for each
 * opcode -- indeces correspond to opcode numbers, integers stored under those
 * indeces represent the number of operands for that specific opcode. It will be
 * used for the semantic analysis and error checking.
 *
 * For example, PUSH_CONST is the second opcode after HALT and thus indexed [1].
 * It takes exactly one operand -- the value it is supposed to push onto the
 * data stack.
 *
 * Only the opcodes that require one or more operands are given the value in the
 * array explicitly -- it is because every index in an array of integers that 
 * was initialized using just the size (like we did with "number_of_operands") 
 * would be set to 0 by default anyways.
 */
int number_of_operands[128];
void init_number_of_operands()
{
    number_of_operands[1] = 1; // PUSH_CONST
    number_of_operands[3] = 1; // BINARY_OP
}



// FUNCTION DECLARATIONS
void usage()
{
    std::string u = 
        "Usage details:\n "
        "\t(Windows) CMD:\n "
        "\t~$ sasm.exe <infile path>/<filename>.sasm <outfile path>/"
                                                            "<filename>.svmb\n "
        "\n "
        "\t(Mac OS / Linux) Terminal:\n "
        "\t~$ vm <infile path>/<filename>.sasm <outfile path>/<filename>.svmb";
    std::cout << u << "\n";
}


bool syntax_error_detected(std::vector<TOKEN> toks)
{
    int toks_size = toks.size();
    for (int i = 0; i < toks_size; i++)
    {
        TOKEN t = toks[i];
        if (t.type == "error") return true;
    }
    return false;
}


bool semantic_error_detected(std::vector<TOKEN> toks)
{
    int number_of_operands_given = toks.size() - 1;
    int number_of_operands_needed = number_of_operands[ toks[0].int_val ];
    return number_of_operands_needed != number_of_operands_given;
}



int main(int argc, char* argv[])
{
    // PREPARATION
    init_number_of_operands();
    
    
    // CHECK INPUT FILE
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
    
    
    // READ FILE LINE BY LINE
    unsigned int line_count = 1;
    std::vector<unsigned char> buffer;
    std::ifstream ifile(filename);
    std::string line;
    while ( std::getline(ifile, line) )
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
        /*----------------------------- LEX DEBUG -----------------------------/
        std::cout << line << "\n"; 
        for (int i = 0; i < toks.size(); i++)
        {
            TOKEN t = toks[i];
            std::cout << t.type << " "; 
        } std::cout << "\n\n";
        /*--------------------------------------------------------------------*/
        
        
        // RUN SYNTACTIC ANALYSIS
        if ( syntax_error_detected(toks) )
        {
            printf("\nSYNTAX ERROR on line %d.\n", line_count);
            printf("Assembler exited with non-zero return value.\n");
            return 2; // return code 2 = "syntax error"
        }
        
        
        // RUN SEMANTIC ANALYSIS
        if ( semantic_error_detected(toks) )
        {
            printf("\nSEMANTIC ERROR on line %d.\n", line_count);
            printf("Assembler exited with non-zero return value.\n");
            return 3; // return code 3 = "semantic error"
        }
        
        
        // RUN BYTECODE GEN
        if ( !bytecode_gen(toks, &buffer) )
        {
            printf("\nBYTECODE GENERATION ERROR on line %d.\n", line_count);
            printf("Assembler exited with non-zero return value.\n");
            return 4; // return code 4 = "bytecode generation error"
        }
        /*-------------------------- BYTECODE DEBUG ---------------------------/
        printf("Buffer state at line %d\n", line_count);
        for (char c : buffer)
            printf("%d ", c);
        printf("\n\n");
        /*--------------------------------------------------------------------*/
        
        
        line_count++;
    }
    buffer.push_back( (unsigned char)0 ); // append HALT opcode in the end
    ifile.close();
    
    
    // WRITE TO FILE
    std::ofstream ofile;
    
    // check if output filename was specified
    if (argc < 3) ofile.open("a.scpb"); // dump file into the current directory
    else
    {
        // write to path specified
        filename = argv[2]; 
        ofile.open(filename);
    }
    
    // error checking
    if (!ofile)
    {
        printf("\nFAILED TO OPEN file %s.\n", filename);
        printf("Assembler exited with non-zero return value.\n");
        return 5; // return code 5 = "file open error"
    }
    
    // write to file
    for (char c : buffer) ofile << c;
    ofile.close();
    
    
    return 0; // return code 0 = "OK"
}
