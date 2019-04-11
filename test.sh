echo Compiling: Assembler, Virtual Machine
g++ Assembler/sasm.cpp     -o Assembler/compiled/sasm.out
g++ VirtualMachine/svm.cpp -o VirtualMachine/compiled/svm.out


echo Running: $1.scpa
printf "\n"
printf "\n"
Assembler/compiled/sasm.out     Testing/$1.scpa Testing/$1.scpb
VirtualMachine/compiled/svm.exe                 Testing/$1.scpb


printf "\n"
echo Sanitizing: $1.scpb
rm Testing/$1.scpb
