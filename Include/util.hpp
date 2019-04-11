/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "util.hpp" header file. 
 * It is used as a header file for the main most general class, type, and
 * function definitions and is shared between multiple files.
 *
 */
 


#pragma once

#include <iostream>
#include <vector>
#include <fstream>



typedef uint8_t BYTE;



class File
{
    char* filename;
    std::vector<BYTE> contents;

    public:
    File (char* filename) { this->filename = filename; }

    bool exists()
    {
        std::ifstream tester(this->filename);
        return tester.good();
    }

    void read()
    {
        if ( !this->exists() )
        {
            std::cout << "File does not exist.\n";
            return;
        }
        std::ifstream file;
        file.open(this->filename);
        BYTE b;
        while (file >> b) this->contents.push_back(b);
    }

    std::vector<BYTE> get_contents() { return this->contents; }
}; 

