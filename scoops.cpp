/**/



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