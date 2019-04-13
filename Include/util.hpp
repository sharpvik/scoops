/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "util.hpp" header file. 
 * It is used as a header file for the main most general class, type, and
 * function definitions and is shared between multiple files.
 *
 */
 


#pragma once

#include <stdio.h>
#include <iostream>
#include <vector>
#include <fstream>



typedef uint8_t BYTE;



class File
{
    char* filename;
    BYTE* contents;

    public:

    File (char* filename) { this->filename = filename; }

    bool exists()
    {
        std::ifstream tester(this->filename);
        bool e = tester.good();
        tester.close();
        return e;
    }

    unsigned int len()
    {
        std::ifstream tester(this->filename, std::ios::in | std::ios::binary);
        tester.seekg(0, std::ios::end);
        unsigned int len = tester.tellg();
        tester.close();
        return len;
    }

    void read()
    {
        if ( !this->exists() )
        {
            std::cout << "File does not exist.\n";
            return;
        }

        unsigned int len = this->len();
        std::ifstream file(this->filename, std::ios::in | std::ios::binary);
        
        char* buffer = new char[len];
        file.read(buffer, len);
        file.close();

        this->contents = reinterpret_cast<BYTE*>(buffer);
    }

    BYTE* get_contents() { return this->contents; }
}; 

