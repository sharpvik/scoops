/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This "err.hpp" header file contains a function used to check lexer's output
 * for errors. The "error_detected" fucntion is employed by Scoops Assembler's
 * main "sasm.cpp" file.
 *
 */



#pragma once



// ERROR CHECK
bool error_detected(std::vector<TOKEN> toks)
{
    int toks_size = toks.size();
    for (int i = 0; i < toks_size; i++)
    {
        TOKEN t = toks[i];
        if (t.type == "error") return true;
    }
    return false;
}
