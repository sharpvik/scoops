#pragma once



#include <stdio.h>
#include <vector>
#include <stack>
#include <stack>
#include <utility> // std::pair

#include "../Include/util.hpp"
#include "../Include/opcodes.hpp"
#include "../Include/objects.hpp"



#define INSTRUCTION std::vector<BYTE>



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

