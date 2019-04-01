// STD LIB IMPORTS
#include <stdio.h>
#include <string.h>
#include <stack>


// OWN IMPORTS
#include "util.h"
#include "exec.h"



int main(int argc, char* argv[])
{
    // READ FILE
    char* filename;
    char* bytes;
    unsigned int length;
    // check if filename was specified
    if (argc > 1) filename = argv[1];
    else 
    {
        printf("Usage: vm.exe <filename>\n");
        return 1; // return code 1 = "no filename specified"
    }
    // check if file exists
    if ( FileExists(filename) )
    {
        bytes = ReadFile(filename);
        length = strlen(bytes);
    }
    else {
        printf("File does not exist.\n");
        return 404; // return code 404 = "file not found"
    }

    
    // EXECUTION LOOP
    char *ip = &bytes[0];
    while (*ip != 4)
    {
        //printf("%u\n", *ip);
        ip = exec(ip);
    }
    
    
    return 0;
}