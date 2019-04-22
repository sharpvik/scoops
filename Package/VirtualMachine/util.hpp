/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "util.hpp" header file for the Scoops Virtual Machine (SVM).
 * It contains all of the necessary includes for all of the files located in
 * this folder + some definitions.
 *
 */



#pragma once



#include <stdio.h>
#include <iostream>
#include <vector>
#include <stack>
#include <string>
#include <utility> // std::pair

#include "../Include/util.hpp"
#include "../Include/opcodes.hpp"
#include "../Include/objects.hpp"



#define INSTRUCTION std::vector<BYTE>
// BYTE is defined in the "../Includes/util.hpp"