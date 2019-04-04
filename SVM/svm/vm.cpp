/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the main C++ file for the Scoops Virtual Machine (SVM). It makes
 * use of the header (.hpp) files that are located in this directory
 * (<...>/scoops/SVM/src/) to allow for great code readability and easy editing.
 *
 * The main function of SVM is executing bytecode (.svmb) files brought to life
 * by the Compiler (see 'Compiler' folder for details or refer to the general
 * documentation in the 'Doc' folder).
 *
 * Scoops Virtual Machine ...
 *     1. Checks existance of the file that was passed to it using the
 *        "FileExists" function from the "util.hpp" file;
 *     2. Reads the file into an array of bytes -- declared as "char* bytes" 
 *        in the main function -- using the "ReadFile" function 
 *        from the "util.hpp" file;
 *     2. Initialises stacks;
 *     3. Proceeds into the execution loop where it uses the instruction pointer
 *        to move along the "bytes" array, executing opcodes using the "exec"
 *        function from the "exec.hpp" file;
 *     4. Upon successful completion -- namely when the HALT opcode 
 *        is reached -- returns 0.
 *
 *
 * Usage details:
 *     CMD:
 *     ~$ vm.exe <filename>.svmb
 *
 *     (Mac OS / Linux) Terminal:
 *     ~$ vm <filename>.svmb
 *
 * 
 * Return codes and their meanings:
 *     0   = "OK"
 *     1   = "no filename specified"
 *     404 = "file not found"
 *
 */



#include <stdio.h>
#include <string.h>
#include <stack>

#include "util.hpp"
#include "exec.hpp"



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



int main(int argc, char* argv[])
{
    // READ FILE
    char* filename;
    char* bytes;
    unsigned int length;
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
    bytes = ReadFile(filename);
    length = strlen(bytes);
    
    
    // STACKS
    std::stack<OBJECT> data;

    
    // EXECUTION LOOP
    char *ip = bytes;
    while (*ip != HALT)
    {
        //printf("opcode %d found\n", *ip); // opcode detection debug
        ip = exec(ip, &data);
    }
    
    
    return 0; // return code 0 = "OK"
}
