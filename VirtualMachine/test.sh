echo Compiling: Virtual Machine
g++ svm.cpp -o compiled/svm.out


echo Running: $1.scpb
printf "\n"
printf "\n"
compiled/svm.out ../Testing/scpb/$1.scpb
