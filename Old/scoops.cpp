/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "scoops.scp" file -- the main file in this directory. It is
 * equivalent to The Glove of Thanos as it contains the power of all six
 * infinity stones from the "Package" folder:
 *     1. Include;
 *     2. Compiler;
 *     3. Assembly Reader;
 *     4. Assembler;
 *     5. Bytecode Reader;
 *     6. Virtual Machine.
 *
 * Each one of these stones grants you enormous power. Yet, you have to remember
 * that great power comes with great responsibility.
 *
 */



#include <stdio.h>
#include <iostream>



void usage()
{
    printf(
        "\n"
        "\tHelp on Usage -- Scoops:\n"
        "\n"
        "\t| General format:\n"
        "\t|---- scoops -<flag> <source>\n"
        "\n"
        "\t| Flag is optional. Source file name must be provided.\n"
        "\n"
        "\t| Allowed flags:\n"
        "\t|---- [empty] = Run (*.scp / *.scpa / *.scpb)\n"
        "\t|---- [-a]    = Assemble (*.scp -> *.scpa)\n"
        "\t|---- [-c]    = Compile (*.scp / *.scpa -> *.scpb)\n"
        "\n"
        "\t| Allowed file extentions:\n"
        "\t|---- [.scp]  = Scoops Source File\n"
        "\t|---- [.scpa] = Scoops Assembly File\n"
        "\t|---- [.scpb] = Scoops Bytecode File\n"
        "\n"
    );
}



int main(int argc, char* argv[])
{
    char flag;
    char* src;
    
    if (argc > 3 || argc < 2)
    {
        printf("\n\tERROR: WRONG NUMBER OF COMMAND LINE ARGUMENTS.\n");
        usage();
        return 1;
    }
    
    for (int i = 1; i < argc; i++)
    {
        char* arg = argv[i];
        if (arg[0] == '-') flag = arg[1];
        else src = arg;
    }
    
    std::cout << flag << " " << src << "\n";
    
    return 0;
}