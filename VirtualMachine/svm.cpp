/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the main file for the Scoops Virtual Machine (SVM). It makes
 * use of the header (.hpp) files that are located in this directory
 * (<...>/scoops/SVM/svm/) to allow for great code readability and easy editing.
 *
 * The main function of SVM is executing bytecode (.scpb) files brought to life
 * by the Scoops Assembler (SASM). SASM source code can be found in the
 * following directory: (<...>/scoops/SVM/sasm/).
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
#include <vector>
#include <stack>

#include "util.hpp"
#include "exec.hpp"
//#include "ops.hpp"



int main(int argc, char* argv[])
{
    // READ FILE
    char* filename;
    // check if filename was specified
    if (argc < 2)
    {
        usage();
        return 1; // return code 1 = "no filename specified"
    }
    File src{argv[1]};
    // check if file exists
    if ( !src.exists() )
    {
        printf("File does not exist.\n");
        return 404; // return code 404 = "file not found"
    }
    // finally read it
    src.read();
    unsigned int length = src.len();
    src.read();
    BYTE* bytes = src.get_contents();
    
    
    // STACKS AND MEMORY
    std::vector<INSTRUCTION> memory;
    std::stack<ScpObj> data;


    // GENERATE INSTRUCTIONS FROM "bytes"
    INSTRUCTION tmp;
    for (unsigned int i = 0; i < length; i++)
    {
        BYTE _this = bytes[i];
        if (_this == INSTRUCTION_END) 
        {
            memory.push_back(tmp);
            tmp.clear();
        }
        else tmp.push_back(_this);
    }
    if (tmp.size() > 0) memory.push_back(tmp); // just in case
    
    
    // EXECUTION LOOP
    bool running = true;
    bool error = false;
    unsigned int ip = 0;
    while (running && !error)
        exec(running, error, ip, memory, data);
    
    
    return 0; // return code 0 = "OK"
}
