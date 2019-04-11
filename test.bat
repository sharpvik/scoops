@echo off

echo Compiling: Assembler, Virtual Machine
g++ -o Assembler\compiled\sasm.exe Assembler\sasm.cpp
g++ -o VirtualMachine\compiled\svm.exe VirtualMachine\svm.cpp


echo Running: %1.scpa
echo.
echo.
Assembler\compiled\sasm.exe Testing\%1.scpa Testing\%1.scpb
VirtualMachine\compiled\svm.exe Testing\%1.scpb


echo.
echo Sanitizing: %1.scpb
del Testing\%1.scpb

