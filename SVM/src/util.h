/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "util.h" header file for the Scoops Virtual Machine (SVM). 
 * It is used as a header file for function and type definitions that are shared
 * between multiple files in this directory. Thus, you can see many other files
 * including it, while "util.h" itself mostly uses just the standard library.
 *
 */
 


#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <fstream>



// OPCODES
#define HALT        0
#define EMIT_CONST  1
#define PUSH_BLN    2
#define PUSH_CHR    3
#define PUSH_NUM    4



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



// USAGE
void usage()
{
    std::string u = 
        "Usage details:\n "
        "\t(Windows) CMD:\n "
        "\t~$ vm.exe <filename>.svmb\n "
        "\n "
        "\t(Mac OS / Linux) Terminal:\n "
        "\t~$ vm <filename>.svmb";
    std::cout << u << "\n";
}