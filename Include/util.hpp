/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "util.hpp" header file. 
 * It is used as a header file that is shared between multiple files. Thus,
 * you can see many other files including it, while "util.hpp" itself mostly
 * uses just the standard library.
 *
 */
 


#pragma once

#include <fstream>



class File
{
    char* name;
    char* contents;


    public:

    File (char* filename) { this->name = filename; }


    void set_name(char* filename) { this->name = filename; }


    bool exists(char* filename)
    {
        std::ifstream file(filename);
        return file.good();
    }


    void read(char* filename)
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
        this->contents = buffer;
    }


    char* get_contents() { return this->contents; }
};
    

