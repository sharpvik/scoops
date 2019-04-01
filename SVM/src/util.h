#pragma once

#include <stdio.h>
#include <vector>
#include <fstream>



// TYPE DECLARATIONS
typedef uint8_t BYTE;



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
